package indicators

// VWRSI implements the Volume Weighted Relative Strength Index (VWRSI) indicator.
func VWRSI(prices []float64, volumes []float64, period int) []float64 {
	if len(prices) != len(volumes) {
		return nil
	}

	if period <= 0 {
		return nil
	}

	if len(prices) < period {
		return nil
	}

	vwrsis := make([]float64, len(prices))
	gains := make([]float64, len(prices))
	losses := make([]float64, len(prices))

	for i := 1; i < len(prices); i++ {
		delta := prices[i] - prices[i-1]
		if delta > 0 {
			gains[i] = delta * volumes[i]
			losses[i] = 0
		} else {
			gains[i] = 0
			losses[i] = -delta * volumes[i]
		}
	}

	for i := period; i < len(prices); i++ {
		sumGain := 0.0
		sumLoss := 0.0

		for j := i - period + 1; j <= i; j++ {
			sumGain += gains[j]
			sumLoss += losses[j]
		}

		if sumLoss == 0 {
			vwrsis[i] = 100
		} else {
			rs := sumGain / sumLoss
			vwrsis[i] = 100 - (100 / (1 + rs))
		}
	}

	return vwrsis
}
