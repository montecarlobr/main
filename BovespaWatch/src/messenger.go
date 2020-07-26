package main

import (
    "fmt"
    "os"
    "log"
    "time"
    "github.com/subosito/twilio"
    "github.com/vic3lord/stocks"
)

var (
    AccountSid string = "ACfadd593a319189b092b57f19f9a8e8d0"
    AuthToken string = "23e8cea197e1bdc89b38be51d9d4ebc3"
    twilioNumber string = "+17472342017"
    traderNumber string = "+5581982599865"
    traderFullName string = "Rodrigo Alves"
)

func getStockAndPrice() []string {
    argsWithProg := os.Args[1:]
    stockCode := argsWithProg[0]
    targetPrice := argsWithProg[1]

    return []string{stockCode, targetPrice}
}

func buildAckMessage(s string, tP string) string {
    ackMessage := fmt.Sprintf("Registering event for %s at target price R$%s", s, tP)

    return ackMessage
}

func deliverSMS(responseMessage string) {
    params := twilio.MessageParams {
        Body: responseMessage,
    }

    twilioClient := twilio.NewClient(AccountSid, AuthToken, nil)

    s, response, err := twilioClient.Messages.Send(twilioNumber, traderNumber, params)
    if err != nil {
        log.Fatal(s, response, err)
    } else {
        log.Println("Message successfully delivered.")
    }
}

func main() {
    stockAndPrice := getStockAndPrice()
    stockCode := stockAndPrice[0]
    targetPrice := stockAndPrice[1]

    log.Println(buildAckMessage(stockCode, targetPrice))

    yFinanceStockCode := fmt.Sprintf("%s.SA", stockCode)

    stock, err := stocks.GetQuote(yFinanceStockCode)
    if err != nil {
        fmt.Println(err)
    } else {
        price, err := stock.GetPrice()
        if err != nil {
            log.Fatal(price, err)
        } else {
            responseMessage := fmt.Sprintf("Hi %s, %s at R$%.2f on %s", traderFullName, stockCode, price, time.Now().Format(time.RFC1123))

            deliverSMS(responseMessage)
        }
    }
}
