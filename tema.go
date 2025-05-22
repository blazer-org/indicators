package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
)

// TEMA calculates the Triple Exponential Moving Average.
func TEMA(prices []float64, timePeriod int32) []float64 {
	n := len(prices)

	// Edge case: empty prices slice
	if n == 0 {
		return nil
	}

	// TA-Lib's TEMA typically requires timePeriod >= 2.
	// Source: TA-Lib C API documentation for TA_TEMA.
	if timePeriod < 2 {
		return []float64{}
	}

	// Determine minimum length required for TA_TEMA.
	// The lookback for TEMA is `3 * timePeriod - 2`.
	// Output begins at index Lookback. So, len(prices) must be at least Lookback + 1
	// for TA-Lib to produce a non-zero/non-NaN value.
	// So, n >= (3 * timePeriod - 2).
	// cgo-talib.Tema returns 0s for the initial unstable period.
	// Our wrapper ensures that we have at least the minimum number of elements
	// for the first actual TEMA value to be computed.
	requiredLength := 3*timePeriod - 2
	if int32(n) < requiredLength {
		return []float64{}
	}

	// Call cgo-talib's Tema function
	temaValues := cgotalib.Tema(prices, timePeriod)
	return temaValues
}
