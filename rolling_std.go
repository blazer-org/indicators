package indicators

import "math"

func RollingStd(data []float64, window int) []float64 {
	result := make([]float64, len(data))

	if window <= 0 || len(data) < window {
		return result
	}

	// Calculate the rolling standard deviation
	for i := range data {
		if i+1 < window {
			result[i] = math.NaN()
			continue
		}

		sum := 0.0
		sumSq := 0.0
		for j := i - window + 1; j <= i; j++ {
			sum += data[j]
			sumSq += data[j] * data[j]
		}

		mean := sum / float64(window)
		// Sample variance: divide by (window - 1)
		variance := (sumSq - float64(window)*mean*mean) / float64(window-1)
		if variance < 0 {
			variance = 0
		}
		result[i] = math.Sqrt(variance)
	}
	return result
}
