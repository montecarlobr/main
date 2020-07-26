package main

import (
    "fmt"
    "os"
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

func getStockAndPrice() []string {
    argsWithProg := os.Args[1:]
    stockCode := argsWithProg[0]
    targetPrice := argsWithProg[1]

    return []string{stockCode, targetPrice}
}

func buildAckMessage(s string, tP string) string {
    ackMessage := fmt.Sprintf("Successfully registered event for %s at target price R$%s", s, tP)

    return ackMessage
}

func getDatabase() *sql.DB {
    db, err := sql.Open("postgres", "user=bovespawatch password=test dbname=bovespawatch sslmode=disable")

    if err != nil {
        log.Fatal(err)
    }

    return db
}

var (
    db *sql.DB = getDatabase()
)

func main() {
    stockAndPrice := getStockAndPrice()
    stockCode := stockAndPrice[0]
    targetPrice := stockAndPrice[1]

    result, err := db.Exec("INSERT INTO watch_orders (stockCode, targetPrice, traderId) VALUES($1, $2, $3)", stockCode, targetPrice, 1)

    if err != nil {
        log.Fatal(result, err)
    } else {
        log.Println(buildAckMessage(stockCode, targetPrice))
    }
}
