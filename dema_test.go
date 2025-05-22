package indicators

import (
	"testing"

	cgotalib "github.com/blazer-org/cgo-talib"
	"github.com/stretchr/testify/assert"
)

func TestDEMA(t *testing.T) {
	emptyResult := []float64{}
	nilResult := []float64(nil)

	// Test Case 1: Basic valid case
	// DEMA lookback is 2*timePeriod - 1. Output starts after that.
	// cgo-talib returns 0.0 for the initial (2*timePeriod - 2) elements.
	prices1 := []float64{
		81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.84, 83.99, 84.55, 84.36, // 1-10
		85.53, 86.54, 86.89, 87.77, 87.29, 88.07, 88.23, 87.30, 87.62, 86.50, // 11-20
		86.83, 87.02, 86.62, 86.41, 85.97, 86.25, 87.38, 87.62, 87.38, 87.63, // 21-30
		87.03, 86.53, 86.25, // 31-33. Min length for period 5 is 2*5-1 = 9.
	} // Length 33.
	timePeriod1 := int32(5) // Lookback = 2*5-1 = 9. First output at index 9.

	// Use cgo-talib directly to get the expected output for comparison.
	expected1 := cgotalib.Dema(prices1, timePeriod1)
	actual1 := DEMA(prices1, timePeriod1)

	assert.Equal(t, len(expected1), len(actual1), "Test Case 1 Failed: Length mismatch")
	for i := range expected1 {
		assert.InDelta(t, expected1[i], actual1[i], 0.0001, "Test Case 1 Failed: Value mismatch at index %d, expected %f, got %f", i, expected1[i], actual1[i])
	}

	// Test Case 2: Empty input prices slice
	actual2 := DEMA([]float64{}, timePeriod1)
	assert.Equal(t, nilResult, actual2, "Test Case 2 Failed: Empty prices slice should return nil")

	// Test Case 3: timePeriod < 2
	actual3a := DEMA(prices1, 1)
	assert.Equal(t, emptyResult, actual3a, "Test Case 3a Failed: timePeriod = 1 should return empty slice")
	actual3b := DEMA(prices1, 0)
	assert.Equal(t, emptyResult, actual3b, "Test Case 3b Failed: timePeriod = 0 should return empty slice")
	actual3c := DEMA(prices1, -1)
	assert.Equal(t, emptyResult, actual3c, "Test Case 3c Failed: timePeriod < 0 should return empty slice")


	// Test Case 4: Insufficient data length for calculation
	// DEMA lookback is 2*timePeriod - 1.
	// For timePeriod=5, lookback is 9. Minimum length for output is 9.
	// Our wrapper returns empty if len < 2*timePeriod - 1.
	requiredLen4 := (2*timePeriod1 - 1) // 9 for timePeriod1=5
	
	actual4_less := DEMA(prices1[:requiredLen4-1], timePeriod1) // len = 8
	assert.Equal(t, emptyResult, actual4_less, "Test Case 4 Failed: len < (2*timePeriod-1) should return empty slice from wrapper")

	// Test Case 4b: Data length exactly 2*timePeriod - 1 (should produce output from cgo-talib)
	prices4b := prices1[:requiredLen4] // len = 9
	expected4b := cgotalib.Dema(prices4b, timePeriod1)
	actual4b := DEMA(prices4b, timePeriod1)
	assert.Equal(t, len(expected4b), len(actual4b), "Test Case 4b Failed: Length mismatch for exact required length")
	for i := range expected4b {
		assert.InDelta(t, expected4b[i], actual4b[i], 0.0001, "Test Case 4b Failed: Value mismatch for exact required length at index %d", i)
	}


	// Test Case 5: timePeriod = 2 (minimum valid for DEMA logic)
	// Lookback = 2*2-1 = 3. Min length = 3.
	timePeriod5 := int32(2)
	prices5 := []float64{10, 11, 12, 13, 14} // len 5
	expected5 := cgotalib.Dema(prices5, timePeriod5)
	actual5 := DEMA(prices5, timePeriod5)
	assert.Equal(t, len(expected5), len(actual5), "Test Case 5 Failed: Length mismatch for timePeriod=2")
	for i := range expected5 {
		assert.InDelta(t, expected5[i], actual5[i], 0.0001, "Test Case 5 Failed: Value mismatch for timePeriod=2 at index %d", i)
	}
	
	// Test Case 6: Real-world-like data (longer series)
	prices6 := []float64{
		81.59, 81.06, 82.87, 83.00, 83.61, 83.15, 82.84, 83.99, 84.55, 84.36, 
		85.53, 86.54, 86.89, 87.77, 87.29, 88.07, 88.23, 87.30, 87.62, 86.50, 
		86.83, 87.02, 86.62, 86.41, 85.97, 86.25, 87.38, 87.62, 87.38, 87.63, 
		87.03, 86.53, 86.25, 85.00, 85.35, 85.00, 83.87, 84.34, 85.23, 84.82, 
		84.84, 85.00, 84.82, 84.65, 84.30, 84.77, 84.89, 85.51, 86.18, 86.12,
	} // Length 50
	timePeriod6 := int32(10) // Lookback = 2*10-1 = 19. Min length = 19.
	expected6_full := cgotalib.Dema(prices6, timePeriod6)
	actual6_full := DEMA(prices6, timePeriod6)
	assert.Equal(t, len(expected6_full), len(actual6_full), "Test Case 6 Failed: Length mismatch for longer series")
	for i := range expected6_full {
		assert.InDelta(t, expected6_full[i], actual6_full[i], 0.0001, "Test Case 6 Failed: Value mismatch for longer series at index %d", i)
	}
}
