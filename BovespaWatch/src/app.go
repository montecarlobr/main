package main

import (
    "fmt"
    "time"
    "log"
    "strings"
    "database/sql"
    _ "github.com/lib/pq"
    // the actual postgresql driver. The underscore before the library means that we import pq without side effects. Basically, it means Go will only import the library for its initialization. For pq, the initialization registers pq as a driver for the SQL interface.
)

// Models
type Trader struct {
    id int
    fullName string
    emailAddress string
    phoneNumber string
    monthlyQuota int
    createdAt time.Time
}

type WatchOrder struct {
    id int
    stockCode string
    targetPrice float64
    traderId int
    trader Trader
    active bool
    createdAt time.Time
}

type Message struct {
    id int
    traderId int
    trader Trader
    stockCode string
    fullText string
    createdAt time.Time
}

func (wo WatchOrder) createMessage() Message {
    fullText := fmt.Sprintf("Hi %s, %s at R$%.2f on %s", strings.TrimSpace(wo.trader.fullName), strings.TrimSpace(wo.stockCode), wo.targetPrice, wo.createdAt.Format(time.RFC1123))

    message := Message {
        trader: wo.trader,
        stockCode: wo.stockCode,
        fullText: fullText,
    }

    return message
}

func main() {
    fmt.Println("The Bovespa Watch")

    db, err := sql.Open("postgres", "user=bovespawatch password=test dbname=bovespawatch sslmode=disable")

    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM watch_orders WHERE active IS TRUE")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    orders := make([]*WatchOrder, 0)
    for rows.Next() {
        order := new(WatchOrder)

        err := rows.Scan(&order.id, &order.stockCode, &order.targetPrice, &order.traderId, &order.active, &order.createdAt)

        if err != nil {
            log.Fatal(err)
        }

        traderQuery := fmt.Sprintf("SELECT * FROM traders WHERE id = %d", order.traderId)

        rows, err := db.Query(traderQuery)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        for rows.Next() {
            trader := new(Trader)

            err := rows.Scan(&trader.id, &trader.fullName, &trader.emailAddress, &trader.phoneNumber, &trader.monthlyQuota, &trader.createdAt)

            if err != nil {
                log.Fatal(err)
            }

            order.trader = *trader
            orders = append(orders, order)
        }
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    for _, order := range orders {
        message := order.createMessage()
        fmt.Println(message.fullText)
    }
}
