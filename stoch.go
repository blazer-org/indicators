package indicators

import (
	"math"
)

// StochasticResult holds the %K and %D series
type StochasticResult struct {
	StochK       []float64
	StochKSignal []float64
}

// StochasticOscillator calculates the %K and %D series
func StochasticOscillator(high, low, close []float64, window, smoothWindow int, fillNa bool) StochasticResult {
	n := len(close)
	stochK := make([]float64, n)

	for i := 0; i < n; i++ {
		if i+1 < window {
			stochK[i] = math.NaN()
			continue
		}
		lowMin := math.MaxFloat64
		highMax := -math.MaxFloat64
		for j := i + 1 - window; j <= i; j++ {
			if low[j] < lowMin {
				lowMin = low[j]
			}
			if high[j] > highMax {
				highMax = high[j]
			}
		}
		denom := highMax - lowMin
		if denom == 0 {
			stochK[i] = 0
		} else {
			stochK[i] = 100 * (close[i] - lowMin) / denom
		}
	}

	stochKSignal := SMA(stochK, smoothWindow)

	// Optional: fill NaN values
	if fillNa {
		for i := range stochK {
			if math.IsNaN(stochK[i]) {
				stochK[i] = 50
			}
		}
		for i := range stochKSignal {
			if math.IsNaN(stochKSignal[i]) {
				stochKSignal[i] = 50
			}
		}
	}

	return StochasticResult{
		StochK:       stochK,
		StochKSignal: stochKSignal,
	}
}
