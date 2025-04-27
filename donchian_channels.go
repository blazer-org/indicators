package indicators

import "math"

// minInSlice returns the minimum value in a slice
func minInSlice(data []float64) float64 {
	min := math.Inf(1)
	for _, v := range data {
		if !math.IsNaN(v) && v < min {
			min = v
		}
	}
	return min
}

// maxInSlice returns the maximum value in a slice
func maxInSlice(data []float64) float64 {
	max := math.Inf(-1)
	for _, v := range data {
		if !math.IsNaN(v) && v > max {
			max = v
		}
	}
	return max
}

// Donchian calculates the Donchian Channels
func Donchian(high, low []float64, lowerLen, upperLen int) ([]float64, []float64, []float64) {
	n := len(high)
	if n == 0 || len(low) != n {
		return nil, nil, nil
	}

	if lowerLen <= 0 {
		lowerLen = 20
	}
	if upperLen <= 0 {
		upperLen = 20
	}

	lower := make([]float64, n)
	upper := make([]float64, n)
	mid := make([]float64, n)

	for i := 0; i < n; i++ {
		if i >= lowerLen-1 {
			window := low[i-lowerLen+1 : i+1]
			lower[i] = minInSlice(window)
		} else {
			lower[i] = math.NaN()
		}
		if i >= upperLen-1 {
			window := high[i-upperLen+1 : i+1]
			upper[i] = maxInSlice(window)
		} else {
			upper[i] = math.NaN()
		}
		if !math.IsNaN(lower[i]) && !math.IsNaN(upper[i]) {
			mid[i] = 0.5 * (lower[i] + upper[i])
		} else {
			mid[i] = math.NaN()
		}
	}

	return lower, upper, mid
}
