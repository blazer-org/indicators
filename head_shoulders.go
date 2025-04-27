package indicators

func HeadShoulders(close []float64, high []float64) []float64 {
	if len(close) != len(high) {
		panic("close and high slices must have the same length")
	}

	headShoulders := make([]float64, len(close))
	for i := 4; i < len(close); i++ {
		if high[i-4] < high[i-3] && high[i-3] > high[i-2] && high[i-2] < high[i-1] && close[i] < close[i-4] {
			headShoulders[i] = 1.0
		} else {
			headShoulders[i] = 0.0
		}
	}

	return headShoulders
}
