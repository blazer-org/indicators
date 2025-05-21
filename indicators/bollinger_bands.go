package indicators

import "math"

// BollingerBandsResult holds the Upper Band, Middle Band, and Lower Band.
type BollingerBandsResult struct {
	UpperBand  []float64
	MiddleBand []float64
	LowerBand  []float64
}

// BollingerBands calculates the Bollinger Bands for a given dataset.
// It uses SMA for the middle band and RollingStd for the standard deviation.
func BollingerBands(data []float64, period int, numStdDev float64) BollingerBandsResult {
	dataLen := len(data)
	result := BollingerBandsResult{
		UpperBand:  make([]float64, dataLen),
		MiddleBand: make([]float64, dataLen),
		LowerBand:  make([]float64, dataLen),
	}

	// Handle edge cases: invalid period, numStdDev, or empty data
	if period <= 0 || numStdDev <= 0 || dataLen == 0 {
		// Return zero-filled result, which is already the case
		return result
	}

	// 1. Calculate the Middle Band (SMA)
	// SMA is expected to return a slice of length dataLen, with leading zeros if len(data) < period.
	middleBand := SMA(data, period)
	copy(result.MiddleBand, middleBand) // Store SMA in the result

	// 2. Calculate the Rolling Standard Deviation
	// RollingStd is expected to return a slice of length dataLen.
	// It uses math.NaN() for initial undefined values.
	rollingStdDev := RollingStd(data, period)

	// 3. & 4. Calculate Upper and Lower Bands
	for i := 0; i < dataLen; i++ {
		// If MiddleBand[i] is 0.0 (from SMA, indicating not enough data yet)
		// or if RollingStd[i] is NaN, then Upper and Lower bands should be 0.0.
		// MiddleBand[i] is already set from SMA.
		if result.MiddleBand[i] == 0.0 || math.IsNaN(rollingStdDev[i]) {
			result.UpperBand[i] = 0.0
			result.LowerBand[i] = 0.0
			// result.MiddleBand[i] is already 0.0 if SMA returned 0.0
		} else {
			stdDevVal := numStdDev * rollingStdDev[i]
			result.UpperBand[i] = result.MiddleBand[i] + stdDevVal
			result.LowerBand[i] = result.MiddleBand[i] - stdDevVal
		}
	}

	return result
}
