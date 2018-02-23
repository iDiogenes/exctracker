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
			if percentGain := calculatePercentGain(exchanges[i].bid, exchanges[index].ask); percentGain > pc && exchanges[index].name != exchanges[i].name {
				var message string
				message = "Buy BTC on " + exchanges[i].name + " for " + strconv.Itoa(int(exchanges[i].bid)) +
					" and sell on " + exchanges[index].name + " for " + strconv.Itoa(int(exchanges[index].ask)) +
					" potential opportunity of ~" + strconv.FormatFloat(percentGain, 'f', 3, 64) + " percent"
				log.Println(message)
			}
		}
	}
}

func calculatePercentGain(bid, ask float64) float64 {
	percent := (ask - bid) / bid * 100
	return percent
}
