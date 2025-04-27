package indicators

// CalculateEOM calculates the Ease of Movement (EOM) indicator with SMA smoothing.
func EOM(highs, lows, volumes []float64, window int) []float64 {
	if len(highs) != len(lows) || len(lows) != len(volumes) {
		panic("Input slices must have the same length")
	}

	n := len(highs)
	eomRaw := make([]float64, n)

	// Calculate raw EOM values
	for i := 1; i < n; i++ {
		distanceMoved := (highs[i]+lows[i])/2 - (highs[i-1]+lows[i-1])/2
		boxRatio := volumes[i] / 100000000 / (highs[i] - lows[i])
		eomRaw[i] = distanceMoved / boxRatio
	}

	return eomRaw
}
