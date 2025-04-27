package indicators

// KVOResult holds the calculated KVO and KVO signal values.
type KVOResult struct {
	KVO       []float64
	KVOSignal []float64
}

// KVO calculates the Klinger Volume Oscillator (KVO) and its signal line.
func KVO(high, low, close, volume []float64) KVOResult {
	length := len(close)
	if len(high) != length || len(low) != length || len(volume) != length {
		panic("Input slices must have the same length")
	}

	hlc3 := make([]float64, length)
	hlc3Diff := make([]float64, length)
	xtrend := make([]float64, length)
	kvo := make([]float64, length)

	// Calculate HLC3 (High + Low + Close) / 3
	for i := 0; i < length; i++ {
		hlc3[i] = (high[i] + low[i] + close[i]) / 3
	}

	// Calculate HLC3 difference
	for i := 1; i < length; i++ {
		hlc3Diff[i] = hlc3[i] - hlc3[i-1]
	}

	// Calculate xtrend
	for i := 0; i < length; i++ {
		if hlc3Diff[i] > 0 {
			xtrend[i] = volume[i] * 100
		} else {
			xtrend[i] = -volume[i] * 100
		}
	}

	// Calculate xfast (EMA of xtrend with fastSpan)
	xfast := EMA(xtrend, 34)

	// Calculate xslow (EMA of xtrend with slowSpan)
	xslow := EMA(xtrend, 55)

	// Calculate KVO (xfast - xslow)
	for i := 0; i < length; i++ {
		kvo[i] = xfast[i] - xslow[i]
	}

	// Calculate KVO signal line (EMA of KVO with signalSpan)
	kvoSignal := EMA(kvo, 13)

	return KVOResult{
		KVO:       kvo,
		KVOSignal: kvoSignal,
	}
}
