package indicators

// InstBlockTrade calculates the institutional block trade signal.
// It first calculates the 50 SMA of the volume.
// Then it checks if volume exceeds the 50 SMA times 5 and Close is less than Open.
// If both conditions are met, it returns 1.0. Otherwise, it returns 0.0.
func InstBlockTrade(open []float64, close []float64, volume []float64) []float64 {
	// Calculate the 50 SMA of the volume
	volSMA := SMA(volume, 50)

	// Initialize the result slice
	result := make([]float64, len(open))

	// Loop through the data and calculate the signal
	for i := 0; i < len(open); i++ {
		if i < 50 {
			result[i] = 0.0 // Not enough data for SMA
			continue
		}
		if volume[i] > volSMA[i]*5 && close[i] < open[i] {
			result[i] = 1.0
		} else {
			result[i] = 0.0
		}
	}

	return result
}
