package indicators

// CalculateZScore calculates the Z-Score for a given time series.
func ZScore(data []float64, window int) []float64 {
	std := RollingStd(data, window)
	sma := SMA(data, window)

	zScore := make([]float64, len(data))

	// Calculate Z-Score
	for i := range data {
		if std[i] == 0 {
			zScore[i] = 0 // Avoid division by zero
		} else {
			zScore[i] = (data[i] - sma[i]) / std[i]
		}
	}

	return zScore
}
