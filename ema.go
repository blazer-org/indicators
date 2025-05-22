package indicators

import cgotalib "github.com/blazer-org/cgo-talib"

func EMA(prices []float64, timePeriod int32) []float64 {
	n := len(prices)
	if n == 0 {
		return nil
	}

	// TA-Lib's EMA function expects timePeriod to be at least 2.
	// See: https://mrjbq7.github.io/ta-lib/func_groups/overlap_studies.html (TA_EMA)
	// "optInTimePeriod: From 2 to 100000"
	// cgo-talib might panic or return unexpected results for timePeriod < 2.
	if timePeriod < 2 {
		return []float64{}
	}

	// cgo-talib.Ema can panic if timePeriod > n.
	// It's better to handle this case explicitly.
	if int32(n) < timePeriod {
		return []float64{}
	}

	// The first few values of the output from cgo-talib.Ema will be 0.0 (not NaN)
	// if there's not enough data for the period.
	// For EMA, the first (timePeriod - 1) values are effectively 0.0.
	return cgotalib.Ema(prices, timePeriod)
}
