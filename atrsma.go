package indicators

import (
	"math"
)

// ATRSMA implements the Average True Range (ATR) using a Simple Moving Average (SMA) smoothing.
func ATRSMA(highs, lows, closes []float64, period int) []float64 {
	n := len(highs)
	if len(lows) != n || len(closes) != n || period <= 0 {
		return nil
	}

	trueRanges := make([]float64, n)

	for i := 0; i < n; i++ {
		highLow := highs[i] - lows[i]
		highClose := math.NaN()
		lowClose := math.NaN()

		if i > 0 {
			highClose = math.Abs(highs[i] - closes[i-1])
			lowClose = math.Abs(lows[i] - closes[i-1])
			trueRanges[i] = math.Max(highLow, math.Max(highClose, lowClose))
		} else {
			trueRanges[i] = highLow // No previous close, so just use high - low
		}
	}

	// Compute simple moving average over trueRanges with padding
	atr := make([]float64, n)
	for i := 0; i < n; i++ {
		if i < period-1 {
			atr[i] = math.NaN() // not enough data
			continue
		}
		sum := 0.0
		for j := i - period + 1; j <= i; j++ {
			sum += trueRanges[j]
		}
		atr[i] = sum / float64(period)
	}

	return atr
}
