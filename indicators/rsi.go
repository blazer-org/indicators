package indicators

import "math"

// RSI calculates the Relative Strength Index.
func RSI(data []float64, period int) []float64 {
	rsiValues := make([]float64, len(data)) // Initialized to 0.0 by default

	if period <= 0 || len(data) < period {
		// Return a slice of zeros, which is already the state of rsiValues
		return rsiValues
	}

	gains := make([]float64, len(data))
	losses := make([]float64, len(data))

	// Calculate price changes, gains, and losses
	// Price changes start from data[1] - data[0], so gains/losses are stored from index 1.
	for i := 1; i < len(data); i++ {
		change := data[i] - data[i-1]
		if change > 0 {
			gains[i] = change
			losses[i] = 0.0
		} else if change < 0 {
			gains[i] = 0.0
			losses[i] = math.Abs(change)
		} else {
			gains[i] = 0.0
			losses[i] = 0.0
		}
	}

	// Calculate initial average gain and loss for the first 'period'
	var sumGains float64
	var sumLosses float64
	// The first 'period' price changes are from gains[1]/losses[1] to gains[period]/losses[period]
	for i := 1; i <= period; i++ {
		sumGains += gains[i]
		sumLosses += losses[i]
	}

	avgGain := sumGains / float64(period)
	avgLoss := sumLosses / float64(period)

	// Calculate first RSI value at index 'period'
	if avgLoss == 0 {
		// If avgLoss is 0, RSI is 100. If avgGain is also 0, some sources say 50, others 0 or 100.
		// Sticking to 100 if avgLoss is 0 as per primary instruction.
		// If both avgGain and avgLoss are zero (no price change over the period), RSI is often considered 50.
		if avgGain == 0 {
			rsiValues[period] = 50.0 
		} else {
			rsiValues[period] = 100.0
		}
	} else {
		rs := avgGain / avgLoss
		rsiValues[period] = 100 - (100 / (1 + rs))
	}

	// Calculate subsequent RSI values using smoothed average
	for i := period + 1; i < len(data); i++ {
		avgGain = ((avgGain * float64(period-1)) + gains[i]) / float64(period)
		avgLoss = ((avgLoss * float64(period-1)) + losses[i]) / float64(period)

		if avgLoss == 0 {
			if avgGain == 0 {
				rsiValues[i] = 50.0
			} else {
				rsiValues[i] = 100.0
			}
		} else {
			rs := avgGain / avgLoss
			rsiValues[i] = 100 - (100 / (1 + rs))
		}
	}

	return rsiValues
}
