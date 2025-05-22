package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
)

// ADX calculates the Average Directional Movement Index.
func ADX(high, low, close []float64, timePeriod int32) []float64 {
	n := len(close)

	// Edge case: empty or mismatched length slices
	if n == 0 || len(high) != n || len(low) != n {
		return nil
	}

	// TA-Lib's ADX typically requires timePeriod >= 2.
	// Source: TA-Lib C API documentation for TA_ADX.
	if timePeriod < 2 {
		return []float64{}
	}

	// Determine minimum length required for TA_ADX.
	// The lookback for TA_ADX is 2 * timePeriod - 1.
	// This means the first valid ADX value will be at index (2 * timePeriod - 2)
	// if TA-Lib follows 0-based indexing for output start after lookback.
	// More precisely, output starts at index (TA_ADX_Lookback(timePeriod)).
	// TA_ADX_Lookback = TA_DX_Lookback(timePeriod) + TA_SMA_Lookback(timePeriod)
	// TA_DX_Lookback = TA_ATR_Lookback(timePeriod) = timePeriod -1 (+ DI lookback)
	// TA_ATR_Lookback = timePeriod -1. So DX lookback is complex.
	// TA-Lib's ADX requires at least `2 * timePeriod - 1` data points to produce the first output.
	// Some sources say `2 * timePeriod` or more for stable values.
	// `cgo-talib` will return 0s for the initial unstable period.
	// We'll set a minimum practical length for our wrapper to avoid calling into cgo-talib with too few points
	// that might lead to all zeros or very unstable initial non-zero values.
	// The first ADX value is typically at index (2*timePeriod - 1) when 0-indexed, meaning 2*timePeriod elements are needed.
	// Or, if the lookback is L, then index L is the first value. So N > L elements.
	// Lookback of ADX is (2*timePeriod - 2). Output begins at index (2*timePeriod - 2).
	// Thus, input slice must have at least (2*timePeriod - 2) + 1 = (2*timePeriod - 1) elements.
	// However, cgotalib.Adx seems to require one more element, possibly due to internal array indexing.
	// Using 2 * timePeriod as a safe minimum length before calling into cgo-talib.
	// If len(close) is 2*timePeriod-1, it panics with index [2*timePeriod-1] on length 2*timePeriod-1.
	// So, length must be at least 2*timePeriod.
	requiredLength := 2 * timePeriod
	if int32(n) < requiredLength {
		return []float64{}
	}

	// Call cgo-talib's Adx function
	adxValues := cgotalib.Adx(high, low, close, timePeriod)
	return adxValues
}
