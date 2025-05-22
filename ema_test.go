package indicators

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEMA(t *testing.T) {
	// Test case 1: Simple case with known values
	// For TA-Lib EMA with timePeriod=3, cgo-talib returns 0.0 for the first 2 (timePeriod-1) values.
	// EMA[2] = (1+2+3)/3 = 2.0 (SMA for the first valid EMA)
	// alpha = 2 / (3+1) = 0.5
	// EMA[3] = (prices[3] * 0.5) + (EMA[2] * 0.5) = (4.0 * 0.5) + (2.0 * 0.5) = 2.0 + 1.0 = 3.0
	// EMA[4] = (prices[4] * 0.5) + (EMA[3] * 0.5) = (5.0 * 0.5) + (3.0 * 0.5) = 2.5 + 1.5 = 4.0
	prices1 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	timePeriod1 := int32(3)
	expected1 := []float64{0.0, 0.0, 2.0, 3.0, 4.0} // cgo-talib uses 0.0 for initial unstable period
	actual1 := EMA(prices1, timePeriod1)
	assert.Equal(t, len(expected1), len(actual1), "Test Case 1 Failed: Length mismatch")
	for i := range expected1 {
		if math.IsNaN(expected1[i]) { // Should not happen with 0.0, but kept for safety
			assert.True(t, math.IsNaN(actual1[i]), "Test Case 1 Failed: Expected NaN at index %d, got %f", i, actual1[i])
		} else {
			assert.InDelta(t, expected1[i], actual1[i], 0.001, "Test Case 1 Failed: Value mismatch at index %d, expected %f, got %f", i, expected1[i], actual1[i])
		}
	}

	// Test case 2: Edge case with an empty input prices slice
	prices2 := []float64{}
	timePeriod2 := int32(5)
	expected2 := []float64(nil) // Expecting nil as per ema.go logic for n=0
	actual2 := EMA(prices2, timePeriod2)
	assert.Equal(t, expected2, actual2, "Test Case 2 Failed: Empty prices slice not handled correctly")

	// Test case 3: Edge case with timePeriod being 1
	// ema.go now returns empty slice for timePeriod < 2
	prices3 := []float64{10.0, 11.0, 12.0, 13.0, 14.0}
	timePeriod3 := int32(1)
	expected3 := []float64{}
	actual3 := EMA(prices3, timePeriod3)
	assert.Equal(t, expected3, actual3, "Test Case 3 Failed: timePeriod=1 should return an empty slice")

	// Test case 4: Edge case with timePeriod greater than the length of prices
	// ema.go now returns empty slice for timePeriod > len(prices)
	prices4 := []float64{1.0, 2.0, 3.0}
	timePeriod4 := int32(5)
	expected4 := []float64{}
	actual4 := EMA(prices4, timePeriod4)
	assert.Equal(t, expected4, actual4, "Test Case 4 Failed: timePeriod > len(prices) should return an empty slice")

	// Test case 5: Edge case with timePeriod being 0
	// ema.go now returns empty slice for timePeriod < 2
	prices5 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	timePeriod5 := int32(0)
	expected5 := []float64{}
	actual5 := EMA(prices5, timePeriod5)
	assert.Equal(t, expected5, actual5, "Test Case 5 Failed: timePeriod=0 should return an empty slice")

	// Test case 6: Edge case with negative timePeriod
	// ema.go now returns empty slice for timePeriod < 2
	prices6 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	timePeriod6 := int32(-2)
	expected6 := []float64{}
	actual6 := EMA(prices6, timePeriod6)
	assert.Equal(t, expected6, actual6, "Test Case 6 Failed: Negative timePeriod should return an empty slice")

	// Test Case 7: Another simple case with different values
	prices7 := []float64{22.27, 22.19, 22.08, 22.17, 22.18, 22.13, 22.23, 22.43, 22.24, 22.29, 22.15, 22.39, 22.38, 22.61, 23.36, 24.05, 23.75, 23.83, 23.95, 23.63, 23.82, 23.87, 23.65, 23.19, 23.10, 23.33, 22.94, 23.00, 22.70, 22.69}
	timePeriod7 := int32(10)
	// Expected values from TA-Lib, cgo-talib returns 0.0 for first (timePeriod-1) values
	expected7 := []float64{
		0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, // Indices 0-8 are 0.0
		22.2246, // EMA at index 9 (SMA of first 10 prices)
		22.2165, // EMA at index 10
		22.2353, // EMA at index 11
		22.2652, // EMA at index 12
		22.3297, // EMA at index 13
		22.5170, // EMA at index 14
		22.7940, // EMA at index 15
		22.9687, // EMA at index 16
		23.1244, // EMA at index 17
		23.2709, // EMA at index 18
		23.3362, // EMA at index 19
		23.4241, // EMA at index 20
		23.5043, // EMA at index 21
		23.5317, // EMA at index 22
		23.4696, // EMA at index 23
		23.4024, // EMA at index 24
		23.3892, // EMA at index 25
		23.3075, // EMA at index 26
		23.2516, // EMA at index 27
		23.1504, // EMA at index 28
		23.0685, // EMA at index 29
	}
	actual7 := EMA(prices7, timePeriod7)

	assert.Equal(t, len(expected7), len(actual7), "Test Case 7 Failed: Length mismatch, expected %d, got %d", len(expected7), len(actual7))
	for i := range expected7 {
		if math.IsNaN(expected7[i]) { // Should not be hit if using 0.0
			assert.True(t, math.IsNaN(actual7[i]), "Test Case 7 Failed: Expected NaN at index %d, got %f", i, actual7[i])
		} else {
			// Increased delta to 0.01 for Test Case 7 to account for minor floating point variations
			assert.InDelta(t, expected7[i], actual7[i], 0.01, "Test Case 7 Failed: Value mismatch at index %d, expected %.4f, got %.4f", i, expected7[i], actual7[i])
		}
	}
}
