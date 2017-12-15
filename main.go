// This program compairs BTC on differnt exchanges

package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/jsgoyette/gemini"
	gdax "github.com/preichenberger/go-gdax"
	twilio "github.com/saintpete/twilio-go"
)

type message struct {
	sentAt  int64
	amount  float64
	message string
}

func main() {
	log.Println("Starting Exchange Tracker.")
	exchangeClients := []interface{}{
		gdaxExc{gdax.NewClient("secret", "key", "passphrase")},
		geminiExc{gemini.New(true, "key", "secret")},
		// krakenExc{krakenapi.New("KEY", "SECRET")}, - Removed kraken, because it kept sending bad data
	}

	for {
		var exchanges exchanges
		for _, exchangeClient := range exchangeClients {
			ec := exchangeClient.(price)
			exchanges = append(exchanges, getPrice(ec))
		}
		arbitrageCheck(exchanges)

		time.Sleep(60 * time.Second)
	}
}

// Check if arbitrage is possible
func arbitrageCheck(exchanges exchanges) {
	exchangeSize := len(exchanges)
	diffPrice, _ := strconv.ParseFloat(os.Getenv("PRICE"), 64)

	// Loop through exchanges and compare ask to bid in all other exchanges
	for index := range exchanges {
		logPrice(exchanges[index].name, exchanges[index].ask)
		for i := 0; i < exchangeSize; i++ {
			if price := (math.Abs(exchanges[index].ask - exchanges[i].bid)); price > diffPrice && exchanges[index].name != exchanges[i].name {
				msg := message{
					amount:  price,
					sentAt:  time.Now().Unix(),
					message: "Buy BTC on " + exchanges[i].name + " and sell on " + exchanges[index].name + " potential opportunity of ~" + strconv.Itoa(int(price)),
				}
				processMessage(&msg)
			}
		}
	}
}

func processMessage(message *message) {
	client := twilio.NewClient(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_TOKEN"), nil)
	msg, err := client.Messages.SendMessage(os.Getenv("TWILIO_FROM"), os.Getenv("TWILIO_TO"), message.message, nil)
	if err != nil {
		log.Println("Issue sending text message", err)
	}
	log.Println(msg.Body)
}

func logPrice(exchange string, ask float64) {
	log.Println("Ask price on " + exchange + " is " + strconv.Itoa(int(ask)))
}
