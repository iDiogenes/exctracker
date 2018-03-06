// This program compairs BTC on differnt exchanges

package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/beldur/kraken-go-api-client"
	"github.com/jsgoyette/gemini"
	"github.com/pdepip/go-binance/binance"
	gdax "github.com/preichenberger/go-gdax"
)

func main() {
	log.Println("Starting Exchange Tracker.")
	logFile, err := os.OpenFile("data/arb_bps.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	exchangeClients := []interface{}{
		gdaxExc{gdax.NewClient("secret", "key", "passphrase")},
		geminiExc{gemini.New(true, "key", "secret")},
		krakenExc{krakenapi.New("KEY", "SECRET")},
		binanceExc{binance.New("", "")},
	}

	for {
		var exchanges exchanges
		for _, exchangeClient := range exchangeClients {
			ec := exchangeClient.(price)
			exchanges = append(exchanges, getPrice(ec))
		}
		arbitrageCheck(exchanges)

		time.Sleep(2 * time.Second)
	}
}
