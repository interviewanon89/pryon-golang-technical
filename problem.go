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

func (s *salesSource) getLatestTransaction() (transactionId string, units int, totalPrice int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Millisecond)

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
		fmt.Printf("time %s -- price %f\n", t.Format(time.RFC3339), p)
	}
}

func histogram(sources []salesSource, histBin time.Duration, histSize time.Duration) (map[time.Time]float32, error) {
	// Initialize variables to store the aggregated data
	histogramData := make(map[time.Time]float32)
	startTime := time.Now()
	endTime := startTime.Add(histSize)
	binTime := startTime

	// Simulate and aggregate sales transactions
	for binTime.Before(endTime) {
		totalPrice := 0
		transactionCount := 0

		// Simulate transactions for each source
		for _, source := range sources {
			_, _, price := source.getLatestTransaction()
			totalPrice += price
			transactionCount++
		}

		// Calculate the average price for the current time bin
		averagePrice := float32(totalPrice) / float32(transactionCount)

		// Store the average price in the histogram data
		histogramData[binTime] = averagePrice

		// Move to the next time bin
		binTime = binTime.Add(histBin)
	}

	return histogramData, nil
}
