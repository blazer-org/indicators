package indicators

import (
	"testing"

	cgotalib "github.com/blazer-org/cgo-talib"
	"github.com/stretchr/testify/assert"
)

func TestKAMA(t *testing.T) {
	emptyResult := []float64{}
	nilResult := []float64(nil)

	// Test Case 1: Basic valid case
	// KAMA lookback is `timePeriod`. Output starts at index `timePeriod`.
	// cgo-talib returns 0.0 for the initial `timePeriod` elements.
	prices1 := []float64{
		81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.84, 83.99, 84.55, 84.36, 
		85.53, 86.54, 86.89, 87.77, 87.29, 88.07, 88.23, 87.30, 87.62, 86.50, 
	} // Length 20
	timePeriod1 := int32(10) // Lookback = 10. First output at index 10.

	// Use cgo-talib directly to get the expected output for comparison.
	expected1 := cgotalib.Kama(prices1, timePeriod1)
	actual1 := KAMA(prices1, timePeriod1)

	assert.Equal(t, len(expected1), len(actual1), "Test Case 1 Failed: Length mismatch")
	for i := range expected1 {
		assert.InDelta(t, expected1[i], actual1[i], 0.0001, "Test Case 1 Failed: Value mismatch at index %d, expected %f, got %f", i, expected1[i], actual1[i])
	}

	// Test Case 2: Empty input prices slice
	actual2 := KAMA([]float64{}, timePeriod1)
	assert.Equal(t, nilResult, actual2, "Test Case 2 Failed: Empty prices slice should return nil")

	// Test Case 3: timePeriod < 2
	actual3a := KAMA(prices1, 1)
	assert.Equal(t, emptyResult, actual3a, "Test Case 3a Failed: timePeriod = 1 should return empty slice")
	actual3b := KAMA(prices1, 0)
	assert.Equal(t, emptyResult, actual3b, "Test Case 3b Failed: timePeriod = 0 should return empty slice")
	actual3c := KAMA(prices1, -1)
	assert.Equal(t, emptyResult, actual3c, "Test Case 3c Failed: timePeriod < 0 should return empty slice")


	// Test Case 4: Insufficient data length for calculation
	// KAMA needs `timePeriod + 1` elements for first output. Our wrapper returns empty if len <= timePeriod.
	// For timePeriod=10.
	// Test with length = timePeriod (10)
	actual4_equal := KAMA(prices1[:timePeriod1], timePeriod1)
	assert.Equal(t, emptyResult, actual4_equal, "Test Case 4 Failed: len == timePeriod should return empty slice from wrapper")
	
	// Test with length < timePeriod (9)
	if timePeriod1 > 1 { 
		actual4_less := KAMA(prices1[:timePeriod1-1], timePeriod1)
		assert.Equal(t, emptyResult, actual4_less, "Test Case 4 Failed: len < timePeriod should return empty slice from wrapper")
	}

	// Test Case 5: Data length just enough for one output value (len = timePeriod + 1)
	prices5 := prices1[:timePeriod1+1] // Length 11 for timePeriod 10
	expected5 := cgotalib.Kama(prices5, timePeriod1)
	actual5 := KAMA(prices5, timePeriod1)
	assert.Equal(t, len(expected5), len(actual5), "Test Case 5 Failed: Length mismatch for exact required length")
	if len(expected5) > 0 { 
		// The first non-zero/non-NaN value is at index timePeriod
		assert.InDelta(t, expected5[len(expected5)-1], actual5[len(actual5)-1], 0.0001, "Test Case 5 Failed: Value mismatch for exact required length at last index")
	}

	// Test Case 6: timePeriod = 2 (minimum valid for KAMA logic)
	// Lookback = 2. Min length for output from wrapper is 3.
	timePeriod6 := int32(2)
	prices6 := []float64{10, 11, 12, 13, 14} // len 5
	expected6 := cgotalib.Kama(prices6, timePeriod6)
	actual6 := KAMA(prices6, timePeriod6)
	assert.Equal(t, len(expected6), len(actual6), "Test Case 6 Failed: Length mismatch for timePeriod=2")
	for i := range expected6 {
		assert.InDelta(t, expected6[i], actual6[i], 0.0001, "Test Case 6 Failed: Value mismatch for timePeriod=2 at index %d", i)
	}
	
	// Test Case 7: Real-world-like data (longer series)
	prices7 := []float64{
		81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.84, 83.99, 84.55, 84.36, 
		85.53, 86.54, 86.89, 87.77, 87.29, 88.07, 88.23, 87.30, 87.62, 86.50, 
		86.83, 87.02, 86.62, 86.41, 85.97, 86.25, 87.38, 87.62, 87.38, 87.63, 
		87.03, 86.53, 86.25, 85.00, 85.35, 85.00, 83.87, 84.34, 85.23, 84.82, 
		84.84, 85.00, 84.82, 84.65, 84.30, 84.77, 84.89, 85.51, 86.18, 86.12, // Length 50
	} 
	timePeriod7 := int32(14) // Lookback = 14. Min length for output from wrapper is 15.
	expected7_full := cgotalib.Kama(prices7, timePeriod7)
	actual7_full := KAMA(prices7, timePeriod7)
	assert.Equal(t, len(expected7_full), len(actual7_full), "Test Case 7 Failed: Length mismatch for longer series")
	for i := range expected7_full {
		assert.InDelta(t, expected7_full[i], actual7_full[i], 0.0001, "Test Case 7 Failed: Value mismatch for longer series at index %d", i)
	}
}
