package main

import (
	"log"
	"os"
	"strconv"
)

// Check if arbitrage is possible
func arbitrageCheck(exchanges exchanges) {
	exchangeSize := len(exchanges)
	pc, _ := strconv.ParseFloat(os.Getenv("PC"), 64)

	// Loop through exchanges and compare ask to bid in all other exchanges
	for index := range exchanges {
		for i := 0; i < exchangeSize; i++ {
			if percentGain := calculatePercentGain(exchanges[index].ask, exchanges[i].bid); percentGain > pc && exchanges[index].name != exchanges[i].name {
				var message string
				message = "Buy BTC on " + exchanges[i].name + " for " + strconv.Itoa(int(exchanges[i].ask)) +
					" and sell on " + exchanges[index].name + " for " + strconv.Itoa(int(exchanges[index].bid)) +
					" potential opportunity of ~" + strconv.FormatFloat(percentGain, 'f', 3, 64) + " percent"
				log.Println(message)
			}
		}
	}
}

func calculatePercentGain(ask, bid float64) float64 {
	percent := (bid*(1-0.0025) - ask) / (2 * ask) * 100
	log.Println(percent)
	return percent
}
