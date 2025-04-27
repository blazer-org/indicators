package indicators

func VWRSI(rsi []float64, volume []float64) []float64 {
	vwRsi := make([]float64, len(rsi))

	if len(rsi) != len(volume) {
		return vwRsi // Ensure rsi and volume have the same length
	}

	for i := range rsi {
		vwRsi[i] = rsi[i] * volume[i]
	}

	return vwRsi
}
