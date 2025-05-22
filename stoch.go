package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
	// "math" // math import will be removed as NaN handling is no longer manual
)

// StochasticResult holds the %K and %D series
type StochasticResult struct {
	StochK       []float64
	StochKSignal []float64
}

// StochasticOscillator calculates the %K and %D series using cgo-talib's Stochf.
// fastDMAType is set to 0 (SMA) for the FastD calculation.
// The fillNa parameter and its logic have been removed; cgo-talib's output (0s for initial unstable periods) is used directly.
func StochasticOscillator(high, low, close []float64, window, smoothWindow int) StochasticResult {
	n := len(close)
	emptyResult := StochasticResult{StochK: []float64{}, StochKSignal: []float64{}}

	// Check for empty or mismatched length input slices
	if n == 0 || len(high) != n || len(low) != n {
		return emptyResult
	}

	// Validate period parameters for TA-Lib.
	// For StochF, fastK_Period (window) and fastD_Period (smoothWindow) must be >= 1.
	// Source: TA-Lib C API documentation for TA_STOCHF.
	if window < 1 || smoothWindow < 1 {
		return emptyResult
	}

	// Determine minimum length required for TA_STOCHF based on its lookback.
	// TA_STOCHF_Lookback = (fastK_Period - 1) + (fastD_Period - 1).
	// Minimum number of elements needed for any output is lookback + 1.
	// So, n >= (window - 1) + (smoothWindow - 1) + 1  => n >= window + smoothWindow - 1.
	requiredLength := (window - 1) + (smoothWindow - 1) + 1
	if n < requiredLength {
		return emptyResult
	}

	fastKPeriod := int32(window)
	fastDPeriod := int32(smoothWindow)
	// Use int32(0) for SMA, as TA-Lib's MA_Type enum typically starts with SMA = 0.
	fastDMAType := int32(0) 

	// Call TA-Lib's Stochastic Fast function (Stochf with lowercase 'f')
	outFastK, outFastD := cgotalib.Stochf(high, low, close, fastKPeriod, fastDPeriod, fastDMAType)

	return StochasticResult{
		StochK:       outFastK,
		StochKSignal: outFastD,
	}
}
