package indicators

import (
	"testing"

	cgotalib "github.com/blazer-org/cgo-talib"
	"github.com/stretchr/testify/assert"
)

func TestMACD(t *testing.T) {
	emptyResult := MACDResult{MACD: []float64{}, MACDSignal: []float64{}, MACDHist: []float64{}}

	// Test Case 1: Basic valid case with default TA-Lib periods
	prices1 := []float64{
		81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.84, 83.99, 84.55, 84.36, // 1-10
		85.53, 86.54, 86.89, 87.77, 87.29, 88.07, 88.23, 87.30, 87.62, 86.50, // 11-20
		86.83, 87.02, 86.62, 86.41, 85.97, 86.25, 87.38, 87.62, 87.38, 87.63, // 21-30
		87.03, 86.53, 86.25, 85.00, 85.35, 85.00, 83.87, 84.34, 85.23, 84.82, // 31-40
		84.84, 85.00, 84.82, 84.65, 84.30, 84.77, 84.89, 85.51, 86.18, 86.12, // 41-50
	}
	fastPeriod1 := int32(12)
	slowPeriod1 := int32(26)
	signalPeriod1 := int32(9)

	// Use cgo-talib directly to get the expected output for comparison.
	expectedMACD1, expectedMACDSignal1, expectedMACDHist1 := cgotalib.Macd(prices1, fastPeriod1, slowPeriod1, signalPeriod1)
	actual1 := MACD(prices1, fastPeriod1, slowPeriod1, signalPeriod1)

	assert.Equal(t, len(expectedMACD1), len(actual1.MACD), "Test Case 1 MACD Failed: Length mismatch")
	assert.Equal(t, len(expectedMACDSignal1), len(actual1.MACDSignal), "Test Case 1 MACDSignal Failed: Length mismatch")
	assert.Equal(t, len(expectedMACDHist1), len(actual1.MACDHist), "Test Case 1 MACDHist Failed: Length mismatch")

	for i := range expectedMACD1 {
		assert.InDelta(t, expectedMACD1[i], actual1.MACD[i], 0.0001, "Test Case 1 MACD Failed: Value mismatch at index %d", i)
	}
	for i := range expectedMACDSignal1 {
		assert.InDelta(t, expectedMACDSignal1[i], actual1.MACDSignal[i], 0.0001, "Test Case 1 MACDSignal Failed: Value mismatch at index %d", i)
	}
	for i := range expectedMACDHist1 {
		assert.InDelta(t, expectedMACDHist1[i], actual1.MACDHist[i], 0.0001, "Test Case 1 MACDHist Failed: Value mismatch at index %d", i)
	}

	// Test Case 2: Empty input prices slice
	actual2 := MACD([]float64{}, fastPeriod1, slowPeriod1, signalPeriod1)
	assert.Equal(t, emptyResult, actual2, "Test Case 2 Failed: Empty prices slice not handled correctly")

	// Test Case 3: Invalid fastPeriod (<1)
	actual3 := MACD(prices1, 0, slowPeriod1, signalPeriod1)
	assert.Equal(t, emptyResult, actual3, "Test Case 3 Failed: fastPeriod < 1 not handled")

	// Test Case 4: Invalid slowPeriod (<1)
	actual4 := MACD(prices1, fastPeriod1, 0, signalPeriod1)
	assert.Equal(t, emptyResult, actual4, "Test Case 4 Failed: slowPeriod < 1 not handled")

	// Test Case 5: Invalid signalPeriod (<1)
	actual5 := MACD(prices1, fastPeriod1, slowPeriod1, 0)
	assert.Equal(t, emptyResult, actual5, "Test Case 5 Failed: signalPeriod < 1 not handled")

	// Test Case 6: fastPeriod >= slowPeriod
	actual6 := MACD(prices1, 26, 12, signalPeriod1) // fast >= slow
	assert.Equal(t, emptyResult, actual6, "Test Case 6 Failed: fastPeriod >= slowPeriod not handled")

	// Test Case 7: Insufficient data length for calculation
	// TA_MACD_Lookback = (slowPeriod - 1) + (signalPeriod - 1). Output starts at index Lookback.
	// For 12, 26, 9: Lookback = (26-1) + (9-1) = 25 + 8 = 33.
	// Minimum data points needed = Lookback + 1 = 34.
	// Our wrapper returns empty if len < Lookback + 1 (i.e., len <= Lookback)
	requiredLookback := (slowPeriod1 - 1) + (signalPeriod1 - 1) // 33
	
	actual7_at_lookback := MACD(prices1[:requiredLookback], fastPeriod1, slowPeriod1, signalPeriod1)
	assert.Equal(t, emptyResult, actual7_at_lookback, "Test Case 7 Failed: len == lookback should return empty from wrapper")
	
	actual7_less_than_lookback := MACD(prices1[:requiredLookback-1], fastPeriod1, slowPeriod1, signalPeriod1)
	assert.Equal(t, emptyResult, actual7_less_than_lookback, "Test Case 7 Failed: len < lookback should return empty from wrapper")


	// Test Case 7b: Data length just enough for one output value
	// Length = requiredLookback + 1 = 34 for (12,26,9)
	prices7b := prices1[:requiredLookback+1] // Length 34
	expectedMACD7b, expectedMACDSignal7b, expectedMACDHist7b := cgotalib.Macd(prices7b, fastPeriod1, slowPeriod1, signalPeriod1)
	actual7b := MACD(prices7b, fastPeriod1, slowPeriod1, signalPeriod1)
	
	assert.Equal(t, len(expectedMACD7b), len(actual7b.MACD), "Test Case 7b MACD Failed: Length mismatch")
	assert.Equal(t, len(expectedMACDSignal7b), len(actual7b.MACDSignal), "Test Case 7b MACDSignal Failed: Length mismatch")
	assert.Equal(t, len(expectedMACDHist7b), len(actual7b.MACDHist), "Test Case 7b MACDHist Failed: Length mismatch")

	if len(expectedMACD7b) > 0 { // cgo-talib might return 1 element or 0 for this specific length
		for i := range expectedMACD7b {
			assert.InDelta(t, expectedMACD7b[i], actual7b.MACD[i], 0.0001, "Test Case 7b MACD Failed: Value mismatch at index %d", i)
		}
		for i := range expectedMACDSignal7b {
			assert.InDelta(t, expectedMACDSignal7b[i], actual7b.MACDSignal[i], 0.0001, "Test Case 7b MACDSignal Failed: Value mismatch at index %d", i)
		}
		for i := range expectedMACDHist7b {
			assert.InDelta(t, expectedMACDHist7b[i], actual7b.MACDHist[i], 0.0001, "Test Case 7b MACDHist Failed: Value mismatch at index %d", i)
		}
	}
}
