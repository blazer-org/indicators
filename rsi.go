package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
)

// RSI calculates the Relative Strength Index.
func RSI(prices []float64, timePeriod int32) []float64 {
	n := len(prices)

	// Edge case: empty prices slice
	if n == 0 {
		return nil
	}

	// TA-Lib's RSI typically requires timePeriod >= 2.
	// Source: TA-Lib C API documentation for TA_RSI.
	if timePeriod < 2 {
		return []float64{}
	}

	// Determine minimum length required for TA_RSI.
	// The first output value of RSI is at index `timePeriod`.
	// This means at least `timePeriod + 1` elements are needed in the input array
	// for the first RSI calculation to be performed.
	// If len(prices) == timePeriod, cgo-talib.Rsi returns all zeros.
	// If len(prices) < timePeriod, cgo-talib.Rsi also returns all zeros (or might panic for very small lengths, though less likely for RSI).
	// To get a meaningful (non-zero, non-NaN) RSI value, len(prices) must be > timePeriod.
	// Our wrapper will return empty if not enough data for any meaningful output.
	if int32(n) <= timePeriod {
		return []float64{}
	}

	// Call cgo-talib's Rsi function
	rsiValues := cgotalib.Rsi(prices, timePeriod)
	return rsiValues
}
