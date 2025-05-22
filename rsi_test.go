package indicators

import (
	"testing"

	cgotalib "github.com/blazer-org/cgo-talib"
	"github.com/stretchr/testify/assert"
)

func TestRSI(t *testing.T) {
	emptyResult := []float64{}
	nilResult := []float64(nil)

	// Test Case 1: Basic valid case with a common period
	prices1 := []float64{
		44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08, // 1-10
		45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64, // 11-20
		46.27, 46.62, 46.82, 46.45, 46.15, 46.00, 46.00, 45.60, 45.07, 45.27, // 21-30
		45.66, // 31
	}
	timePeriod1 := int32(14)

	// Use cgo-talib directly to get the expected output for comparison.
	expected1 := cgotalib.Rsi(prices1, timePeriod1)
	actual1 := RSI(prices1, timePeriod1)

	assert.Equal(t, len(expected1), len(actual1), "Test Case 1 Failed: Length mismatch")
	for i := range expected1 {
		assert.InDelta(t, expected1[i], actual1[i], 0.0001, "Test Case 1 Failed: Value mismatch at index %d, expected %f, got %f", i, expected1[i], actual1[i])
	}

	// Test Case 2: Empty input prices slice
	actual2 := RSI([]float64{}, timePeriod1)
	assert.Equal(t, nilResult, actual2, "Test Case 2 Failed: Empty prices slice should return nil")

	// Test Case 3: timePeriod < 2
	actual3 := RSI(prices1, 1)
	assert.Equal(t, emptyResult, actual3, "Test Case 3 Failed: timePeriod < 2 should return empty slice")
	actual3b := RSI(prices1, 0)
	assert.Equal(t, emptyResult, actual3b, "Test Case 3b Failed: timePeriod = 0 should return empty slice")
	actual3c := RSI(prices1, -1)
	assert.Equal(t, emptyResult, actual3c, "Test Case 3c Failed: timePeriod < 0 should return empty slice")


	// Test Case 4: Insufficient data length for calculation
	// RSI needs `timePeriod + 1` elements. Output starts at index `timePeriod`.
	// If len(prices) <= timePeriod, our wrapper returns empty.
	// Test with length = timePeriod
	actual4_equal := RSI(prices1[:timePeriod1], timePeriod1)
	assert.Equal(t, emptyResult, actual4_equal, "Test Case 4 Failed: len == timePeriod should return empty slice from wrapper")
	
	// Test with length < timePeriod
	if timePeriod1 > 1 { // Ensure we can slice to a shorter length
		actual4_less := RSI(prices1[:timePeriod1-1], timePeriod1)
		assert.Equal(t, emptyResult, actual4_less, "Test Case 4 Failed: len < timePeriod should return empty slice from wrapper")
	}


	// Test Case 5: Data length just enough for one output value (len = timePeriod + 1)
	prices5 := prices1[:timePeriod1+1] // Length 15 for timePeriod 14
	expected5 := cgotalib.Rsi(prices5, timePeriod1)
	actual5 := RSI(prices5, timePeriod1)
	assert.Equal(t, len(expected5), len(actual5), "Test Case 5 Failed: Length mismatch for exact required length")
	if len(expected5) > 0 { // Should have at least one non-zero value if TA-Lib calculated it
		assert.InDelta(t, expected5[len(expected5)-1], actual5[len(actual5)-1], 0.0001, "Test Case 5 Failed: Value mismatch for exact required length at last index")
	}


	// Test Case 6: timePeriod = 2 (minimum valid for RSI logic)
	timePeriod6 := int32(2)
	// Required length for timePeriod=2 is 2+1 = 3 elements.
	prices6 := []float64{10, 11, 12, 13, 14} // len 5
	expected6 := cgotalib.Rsi(prices6, timePeriod6)
	actual6 := RSI(prices6, timePeriod6)
	assert.Equal(t, len(expected6), len(actual6), "Test Case 6 Failed: Length mismatch for timePeriod=2")
	for i := range expected6 {
		assert.InDelta(t, expected6[i], actual6[i], 0.0001, "Test Case 6 Failed: Value mismatch for timePeriod=2 at index %d", i)
	}
	
	// Test Case 7: All input prices are the same (flat line)
	// RSI should be undefined or 0 or 100 depending on implementation details.
	// TA-Lib's RSI typically trends towards 50 in such cases after initial period, or can be 0.
	// For truly flat data from the start, AvgGain and AvgLoss will be 0, leading to RS = undefined.
	// TA-Lib handles division by zero in RS calculation; often results in RSI = 0 or 100.
	// cgo-talib typically returns 0 for initial unstable period.
	prices7 := make([]float64, 30)
	for i := range prices7 {
		prices7[i] = 50.0
	}
	timePeriod7 := int32(14)
	expected7 := cgotalib.Rsi(prices7, timePeriod7) // Let's see what cgo-talib does
	actual7 := RSI(prices7, timePeriod7)
	assert.Equal(t, len(expected7), len(actual7), "Test Case 7 Failed: Length mismatch for flat data")
	for i := range expected7 {
		// For flat data, RSI is often 0 after the initial period.
		// The first few actual RSI values might be 0.0 due to Wilder's smoothing warm-up.
		assert.InDelta(t, expected7[i], actual7[i], 0.0001, "Test Case 7 Failed: Value mismatch for flat data at index %d. Expected: %f, Actual: %f", i, expected7[i], actual7[i])
	}
}
