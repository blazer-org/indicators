package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
)

// DEMA calculates the Double Exponential Moving Average.
func DEMA(prices []float64, timePeriod int32) []float64 {
	n := len(prices)

	// Edge case: empty prices slice
	if n == 0 {
		return nil
	}

	// TA-Lib's DEMA typically requires timePeriod >= 2.
	// Source: TA-Lib C API documentation for TA_DEMA.
	if timePeriod < 2 {
		return []float64{}
	}

	// Determine minimum length required for TA_DEMA.
	// The lookback for DEMA is `2 * timePeriod - 1`.
	// Output begins at index Lookback. So, len(prices) must be at least Lookback + 1
	// for TA-Lib to produce a non-zero/non-NaN value.
	// So, n >= (2 * timePeriod - 1).
	// cgo-talib.Dema returns 0s for the initial unstable period.
	// Our wrapper ensures that we have at least the minimum number of elements
	// for the first actual DEMA value to be computed.
	requiredLength := 2*timePeriod - 1
	if int32(n) < requiredLength {
		return []float64{}
	}

	// Call cgo-talib's Dema function
	demaValues := cgotalib.Dema(prices, timePeriod)
	return demaValues
}
