package main

import (
	"strconv"

	"github.com/beldur/kraken-go-api-client"
	"github.com/jsgoyette/gemini"
	"github.com/pdepip/go-binance/binance"
	gdax "github.com/preichenberger/go-gdax"
)

type price interface {
	get() exchange
}

type exchanges []exchange

type gdaxExc struct {
	client *gdax.Client
}

type geminiExc struct {
	client *gemini.Api
}

type krakenExc struct {
	client *krakenapi.KrakenApi
}

type binanceExc struct {
	client *binance.Binance
}

type exchange struct {
	bid  float64
	ask  float64
	name string
}

// get - Get the bid and ask prices
func (exc gdaxExc) get() exchange {
	ticker, err := exc.client.GetTicker("BTC-USD")
	if err != nil {
		panic(err.Error())
	}
	return exchange{
		bid:  ticker.Bid,
		ask:  ticker.Ask,
		name: "gdax",
	}
}

func (exc geminiExc) get() exchange {
	orderBook, err := exc.client.OrderBook("btcusd", 1, 1)
	if err != nil {
		panic(err)
	}

	return exchange{
		bid:  orderBook.Bids[0].Price,
		ask:  orderBook.Asks[0].Price,
		name: "gemini",
	}
}

func (exc krakenExc) get() exchange {
	ticker, err := exc.client.Ticker(krakenapi.XXBTZUSD)
	if err != nil {
		panic(err)
	}

	ask, _ := strconv.ParseFloat(ticker.XXBTZUSD.Ask[0], 64)
	bid, _ := strconv.ParseFloat(ticker.XXBTZUSD.Bid[0], 64)

	return exchange{
		bid:  bid,
		ask:  ask,
		name: "kraken",
	}
}

func (exc binanceExc) get() exchange {
	query := binance.SymbolQuery{
		Symbol: "BTCUSDT",
	}

	exc.client.GetBookTickers()
	res, err := exc.client.GetLastPrice(query)
	if err != nil {
		panic(err)
	}

	return exchange{
		bid:  res.Price,
		ask:  res.Price,
		name: "binance",
	}
}

func getPrice(exc price) exchange {
	return exc.get()
}
