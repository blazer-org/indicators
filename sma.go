package indicators

import cgotalib "github.com/blazer-org/cgo-talib"

func SMA(prices []float64, timePeriod int32) []float64 {
	n := len(prices)
	if n == 0 {
		return nil
	}

	// TA-Lib's SMA function typically expects timePeriod to be at least 2.
	// Reference (similar to EMA): https://mrjbq7.github.io/ta-lib/func_groups/overlap_studies.html (TA_SMA)
	// "optInTimePeriod: From 2 to 100000"
	// Handle timePeriod < 2 to prevent potential issues with cgo-talib.
	if timePeriod < 2 {
		return []float64{}
	}

	// Handle cases where timePeriod is larger than the number of prices.
	// cgo-talib might panic or return unexpected results in this scenario.
	// TA-Lib functions typically require (timePeriod - 1) elements for the first valid output,
	// so at least `timePeriod` elements are needed in the input for any meaningful calculation.
	if int32(n) < timePeriod {
		return []float64{}
	}

	// cgo-talib.Sma will produce output. The first (timePeriod - 1) elements
	// will represent the unstable period (these are 0.0 for cgo-talib's Sma/Ema).
	return cgotalib.Sma(prices, timePeriod)
}
