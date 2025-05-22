package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
)

// KAMA calculates Kaufman's Adaptive Moving Average.
func KAMA(prices []float64, timePeriod int32) []float64 {
	n := len(prices)

	// Edge case: empty prices slice
	if n == 0 {
		return nil
	}

	// TA-Lib's KAMA typically requires timePeriod >= 2.
	// Source: TA-Lib C API documentation for TA_KAMA.
	if timePeriod < 2 {
		return []float64{}
	}

	// Determine minimum length required for TA_KAMA.
	// The lookback for KAMA is `timePeriod`.
	// The first output value is at index `timePeriod`.
	// This means at least `timePeriod + 1` elements are needed in the input array
	// for the first KAMA calculation to be performed.
	// If len(prices) <= timePeriod, our wrapper returns empty.
	if int32(n) <= timePeriod {
		return []float64{}
	}

	// Call cgo-talib's Kama function
	kamaValues := cgotalib.Kama(prices, timePeriod)
	return kamaValues
}
