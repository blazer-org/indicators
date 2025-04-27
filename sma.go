package indicators

func SMA(data []float64, period int) []float64 {
	// Initialize the result slice
	result := make([]float64, len(data))

	// Calculate the SMA
	for i := 0; i < len(data); i++ {
		if i < period-1 {
			result[i] = 0.0 // Not enough data for SMA
			continue
		}
		sum := 0.0
		for j := 0; j < period; j++ {
			sum += data[i-j]
		}
		result[i] = sum / float64(period)
	}

	return result
}
