package indicators

import "math"

// maxOfThree returns the maximum of three float64 numbers.
func maxOfThree(a, b, c float64) float64 {
	// Simpler way to write max for three numbers
	return math.Max(a, math.Max(b, c))
}

// ATR calculates the Average True Range.
func ATR(high []float64, low []float64, close []float64, period int) []float64 {
	dataLen := len(high)

	// 1. Input Validation
	if len(low) != dataLen || len(close) != dataLen {
		// Panic or return error. For now, return empty slice as per instruction.
		// Consider logging this event in a real application.
		return []float64{}
	}
	if dataLen == 0 {
		return []float64{}
	}
	if period <= 0 {
		return []float64{}
	}
	if dataLen < period {
		// Return a slice of zeros with the same length as high.
		return make([]float64, dataLen)
	}

	// Initialize ATR result slice, all elements default to 0.0
	atr := make([]float64, dataLen)
	tr := make([]float64, dataLen)

	// 2. Calculate True Range (TR)
	// TR[0] = high[0] - low[0]
	if dataLen > 0 { // Should always be true due to earlier checks, but good for safety
		tr[0] = high[0] - low[0]
	}

	for i := 1; i < dataLen; i++ {
		val1 := high[i] - low[i]
		val2 := math.Abs(high[i] - close[i-1])
		val3 := math.Abs(low[i] - close[i-1])
		tr[i] = maxOfThree(val1, val2, val3)
	}

	// 3. Calculate ATR using Wilder's Smoothing
	// Calculate the first ATR value: average of the first 'period' TRs.
	// This value is stored at atr[period-1].
	var sumTR float64
	for i := 0; i < period; i++ {
		sumTR += tr[i]
	}
	if period > 0 { // Should be true due to earlier check
	    atr[period-1] = sumTR / float64(period)
	}


	// Calculate subsequent ATR values using Wilder's smoothing.
	// atr[0] to atr[period-2] remain 0.0.
	for i := period; i < dataLen; i++ {
		atr[i] = (atr[i-1]*float64(period-1) + tr[i]) / float64(period)
	}

	return atr
}
