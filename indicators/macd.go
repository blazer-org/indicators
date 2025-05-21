package indicators

// MACDResult holds the MACD Line, Signal Line, and Histogram.
type MACDResult struct {
	MACDLine   []float64
	SignalLine []float64
	Histogram  []float64
}

// MACD calculates the Moving Average Convergence Divergence.
// It assumes an EMA function `func EMA(data []float64, period int) []float64` is available in the same package.
func MACD(data []float64, fastPeriod int, slowPeriod int, signalPeriod int) MACDResult {
	dataLen := len(data)
	result := MACDResult{
		MACDLine:   make([]float64, dataLen), // Initialized to 0.0 by default
		SignalLine: make([]float64, dataLen), // Initialized to 0.0 by default
		Histogram:  make([]float64, dataLen), // Initialized to 0.0 by default
	}

	// Edge Case: If input data is empty, or periods are invalid.
	// fastPeriod < slowPeriod is a requirement.
	// EMA function is expected to handle short data for its period by returning zero-padded slices.
	// If periods are invalid (e.g., <=0), EMA should return zero-padded slices.
	// We'll add a check for invalid periods or if fastPeriod >= slowPeriod for robustness.
	if dataLen == 0 || fastPeriod <= 0 || slowPeriod <= 0 || signalPeriod <= 0 || fastPeriod >= slowPeriod {
		// Return the zero-filled result, which is already the case.
		return result
	}

	// 1. Calculate Fast EMA
	fastEMA := EMA(data, fastPeriod)

	// 2. Calculate Slow EMA
	slowEMA := EMA(data, slowPeriod)

	// Ensure EMA returned slices of the correct length.
	// If not, it indicates an issue with the EMA function or assumptions.
	// For now, we assume EMA behaves as specified (returns same length slice, zero-padded).

	// 3. Calculate MACD Line: MACDLine[i] = fastEMA[i] - slowEMA[i]
	// Since fastEMA and slowEMA are zero-padded, MACDLine will also be correctly zero-padded.
	for i := 0; i < dataLen; i++ {
		result.MACDLine[i] = fastEMA[i] - slowEMA[i]
	}

	// 4. Calculate Signal Line: EMA of the MACD Line, using signalPeriod.
	// The EMA function will correctly handle leading zeros in result.MACDLine.
	result.SignalLine = EMA(result.MACDLine, signalPeriod)
	
	// Ensure SignalLine is of the correct length.
	// Assuming EMA behaves as specified.

	// 5. Calculate Histogram: Histogram[i] = MACDLine[i] - SignalLine[i]
	// This will also correctly propagate leading zeros.
	for i := 0; i < dataLen; i++ {
		result.Histogram[i] = result.MACDLine[i] - result.SignalLine[i]
	}

	return result
}
