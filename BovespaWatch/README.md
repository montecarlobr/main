# Bovespa Watch
Tool for following stock price updates at the BM&FBovespa.

#### Full Description

This is a simple Go application that will watch changes in the Brazilian stock market - [BM&FBovespa] - in order to notify the user (via the SMS protocol, with the help of [Twilio]) when pre-defined events happen (if they happen).

## Libraries

This project uses several open source softwares to work properly, they are:

* [subosito/twilio]: https://github.com/subosito/twilio
* [lib/pq]: https://github.com/lib/pq
* [vic3lord/stocks]: https://github.com/vic3lord/stocks

## Cost

#### SMS delivery

`$1` for the Twilio number, monthly. `$0.06` per message.
Estimated 120 messages/mo -> `$7.20`.

Little consumption table:

| No. of SMSs | Per day | Total Cost |
| ----------- |:-------:|:----------:|
| 30          | 1.36    | $1.80      |
| 50          | 2.27    | $3.00      |
| 100         | 4.54    | $6.00      |
| 120         | 5.45    | $7.20      |
| 150         | 6.81    | $9.00      |
| 200         | 9.09    | $12.00     |
| 300         | 13.63   | $18.00     |

More at https://www.twilio.com/sms/pricing/br

## Usage

```go
// Receive an SMS in case CSNA3.SA reaches R$ 10,51
$ go run BovespaWatch.go CSNA3 1051
```

## References

* 1 - [Yahoo Stock Quotes in Go]
* 2 - [Going Real-Time with Redis Pub/Sub]
* 3 - [Practical Persistence in Go: SQL Databases]

## Copyright

© 2016 Rodrigo Alves Vieira. All Rights Reserved.

[BM&FBovespa]: http://www.bmfbovespa.com.br/en_us/
[Twilio]: https://www.twilio.com/
[Yahoo Stock Quotes in Go]: https://news.ycombinator.com/item?id=9374373
[Going Real-Time with Redis Pub/Sub]: https://www.toptal.com/go/going-real-time-with-redis-pubsub
[Practical Persistence in Go: SQL Databases]: http://www.alexedwards.net/blog/practical-persistence-sql
