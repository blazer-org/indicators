package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
)

// MACDResult holds the MACD line, signal line, and histogram.
type MACDResult struct {
	MACD       []float64
	MACDSignal []float64
	MACDHist   []float64
}

// MACD calculates the Moving Average Convergence Divergence (MACD).
// Default TA-Lib values are often fastPeriod=12, slowPeriod=26, signalPeriod=9.
func MACD(prices []float64, fastPeriod int32, slowPeriod int32, signalPeriod int32) MACDResult {
	n := len(prices)
	emptyResult := MACDResult{MACD: []float64{}, MACDSignal: []float64{}, MACDHist: []float64{}}

	// Edge case: empty prices slice
	if n == 0 {
		return emptyResult
	}

	// Validate period parameters
	if fastPeriod < 1 || slowPeriod < 1 || signalPeriod < 1 {
		return emptyResult
	}
	if fastPeriod >= slowPeriod {
		return emptyResult
	}

	// Determine minimum length required for TA_MACD.
	// TA_MACD_Lookback = (slowPeriod - 1) + (signalPeriod - 1).
	// Output begins at index Lookback. So, len(prices) must be at least Lookback + 1.
	// n >= (slowPeriod - 1) + (signalPeriod - 1) + 1  => n >= slowPeriod + signalPeriod - 1.
	// cgo-talib's Macd function itself handles short inputs by returning arrays of 0s for the unstable period.
	// However, to align with TA-Lib's concept of a lookback period producing meaningful output,
	// we ensure the length is at least this theoretical minimum for the first full set of values.
	// If TA-Lib C function was directly called, it would specify outBegIdx = (slowPeriod - 1) + (signalPeriod - 1).
	// Let's use this as the minimum required length.
	requiredLength := (slowPeriod - 1) + (signalPeriod - 1)
	if int32(n) <= requiredLength { // Need more than 'lookback' elements for the first output.
		// If n == requiredLength + 1, we get one output.
		// If n == requiredLength, cgo-talib might return all zeros or panic, let's return empty.
		return emptyResult
	}


	// Call cgo-talib's Macd function
	outMACD, outMACDSignal, outMACDHist := cgotalib.Macd(
		prices,
		fastPeriod,
		slowPeriod,
		signalPeriod,
	)

	return MACDResult{
		MACD:       outMACD,
		MACDSignal: outMACDSignal,
		MACDHist:   outMACDHist,
	}
}
