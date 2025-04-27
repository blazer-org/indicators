package indicators

import (
	"math"

	cgotalib "github.com/blazer-org/cgo-talib"
)

type SupertrendResult struct {
	Trend     []float64
	Direction []int
	Long      []float64
	Short     []float64
}

// SuperTrend calculates the Super Trend indicator for a given time series
func Supertrend(high, low, close []float64, length int, multiplier float64) *SupertrendResult {
	n := len(close)
	if n == 0 || len(high) != n || len(low) != n {
		return nil
	}

	// Set default parameters if needed
	if length <= 0 {
		length = 7
	}
	if multiplier <= 0 {
		multiplier = 3.0
	}

	// Initialize output slices
	direction := make([]int, n)
	trend := make([]float64, n)
	long := make([]float64, n)
	short := make([]float64, n)
	for i := 0; i < n; i++ {
		long[i], short[i] = math.NaN(), math.NaN()
	}
	for i := range direction {
		direction[i] = 1 // Start with bullish assumption
	}

	// Compute hl2 (average of high and low)
	hl2 := make([]float64, n)
	for i := 0; i < n; i++ {
		hl2[i] = (high[i] + low[i]) / 2.0
	}

	// Compute ATR
	atr := cgotalib.Atr(high, low, close, int32(length))
	// atr := ATR(high, low, close, length)

	// Compute upperband and lowerband
	upperband := make([]float64, n)
	lowerband := make([]float64, n)
	for i := 0; i < n; i++ {
		upperband[i] = hl2[i] + multiplier*atr[i]
		lowerband[i] = hl2[i] - multiplier*atr[i]
	}

	// Main Supertrend logic
	for i := 1; i < n; i++ {
		if close[i] > upperband[i-1] {
			direction[i] = 1
		} else if close[i] < lowerband[i-1] {
			direction[i] = -1
		} else {
			direction[i] = direction[i-1]
			if direction[i] > 0 && lowerband[i] < lowerband[i-1] {
				lowerband[i] = lowerband[i-1]
			}
			if direction[i] < 0 && upperband[i] > upperband[i-1] {
				upperband[i] = upperband[i-1]
			}
		}

		if direction[i] > 0 {
			trend[i] = lowerband[i]
			long[i] = lowerband[i]
		} else {
			trend[i] = upperband[i]
			short[i] = upperband[i]
		}
	}

	return &SupertrendResult{
		Trend:     trend,
		Direction: direction,
		Long:      long,
		Short:     short,
	}
}
