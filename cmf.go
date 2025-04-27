package indicators

// CMF calculates the Chaikin Money Flow indicator
func CMF(highs, lows, closes, volumes []float64, period int) []float64 {
	n := len(highs)
	if n != len(lows) || n != len(closes) || n != len(volumes) {
		return []float64{}
	}
	if n < period {
		return []float64{}
	}

	mfm := make([]float64, n) // Money Flow Multiplier
	mfv := make([]float64, n) // Money Flow Volume

	for i := 0; i < n; i++ {
		highLowRange := highs[i] - lows[i]
		if highLowRange == 0 {
			mfm[i] = 0 // Avoid division by zero
		} else {
			mfm[i] = ((2*closes[i] - highs[i] - lows[i]) / highLowRange)
		}
		mfv[i] = mfm[i] * volumes[i]
	}

	cmf := make([]float64, n)
	for i := period - 1; i < n; i++ {
		var sumMFV, sumVolume float64
		for j := i - period + 1; j <= i; j++ {
			sumMFV += mfv[j]
			sumVolume += volumes[j]
		}
		if sumVolume == 0 {
			cmf[i] = 0
		} else {
			cmf[i] = sumMFV / sumVolume
		}
	}

	return cmf
}
