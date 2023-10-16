package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	histBinDefault  = 3 * time.Second
	histSizeDefault = 10 * time.Second
)

type salesSource struct {
	name string
}

// Returns the latest sale transaction in units sold and the total price paid for those units.
// NOTE: transactions happen at random, non-deterministic times. The salesSource is session-less and may return the same transaction data if called often enough.
func (s *salesSource) getLatestTransaction() (transactionId string, units int, totalPrice int) {
	// sleep between 20-30 milliseconds to simulate real world delays
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Millisecond)

	// min sale is 1 unit, total price 15
	// max sale is 500 units, total price 12500
	units = 1 + rand.Intn(499)
	totalPrice = units * (15 + rand.Intn(10))
	return
}

func main() {
	h, err := histogram([]salesSource{}, histBinDefault, histSizeDefault)
	if err != nil {
		fmt.Println(fmt.Errorf("Failed with error: %w", err))
		return
	}
	for t, p := range h {
		fmt.Printf("time %s -- price %f/n", t.Format(time.RFC3339), p)
	}
}

// Monitors sale transactions over the given duration from multiple sales sources.
// Computes the average sales price across all given sale sources in real time as the sales occur.
// Returns a histogram of the average unit price over time, aggregated at given time intervals.
//
// For example, if pizza was being sold simultaneously at three sales locations: Seattle, Dallas, and Atlanta,
// the histgram may represent the price fluctuations of a pizza pie sold across the USA at every 5 minute interval for an hour starting now.
//
// sources is a set of sale sources being monitored for transactions
// histBin is the aggregation interval duration, e.g. 5 minutes
// histSize is the total monitoring time duration, e.g. 60 minutes
func histogram(sources []salesSource, histBin time.Duration, histSize time.Duration) (map[time.Time]float32, error) {

	return nil, fmt.Errorf("Not implemented.")
}
