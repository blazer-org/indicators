package indicators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMA(t *testing.T) {
	// Test Case 1: Simple case with known values
	// For timePeriod=3, the first 2 values from cgo-talib are expected to be 0.0.
	// SMA[2] = (1+2+3)/3 = 2.0
	// SMA[3] = (2+3+4)/3 = 3.0
	// SMA[4] = (3+4+5)/3 = 4.0
	prices1 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	timePeriod1 := int32(3)
	expected1 := []float64{0.0, 0.0, 2.0, 3.0, 4.0} // cgo-talib uses 0.0 for initial unstable period
	actual1 := SMA(prices1, timePeriod1)
	assert.Equal(t, len(expected1), len(actual1), "Test Case 1 Failed: Length mismatch")
	for i := range expected1 {
		assert.InDelta(t, expected1[i], actual1[i], 0.001, "Test Case 1 Failed: Value mismatch at index %d, expected %f, got %f", i, expected1[i], actual1[i])
	}

	// Test Case 2: Edge case with an empty input prices slice
	prices2 := []float64{}
	timePeriod2 := int32(5)
	expected2 := []float64(nil) // Expecting nil as per sma.go logic for n=0
	actual2 := SMA(prices2, timePeriod2)
	assert.Equal(t, expected2, actual2, "Test Case 2 Failed: Empty prices slice not handled correctly")

	// Test Case 3: Edge case with timePeriod being 1
	// sma.go returns empty slice for timePeriod < 2
	prices3 := []float64{10.0, 11.0, 12.0, 13.0, 14.0}
	timePeriod3 := int32(1)
	expected3 := []float64{}
	actual3 := SMA(prices3, timePeriod3)
	assert.Equal(t, expected3, actual3, "Test Case 3 Failed: timePeriod=1 should return an empty slice")

	// Test Case 4: Edge case with timePeriod greater than the length of prices
	// sma.go returns empty slice for timePeriod > len(prices)
	prices4 := []float64{1.0, 2.0, 3.0}
	timePeriod4 := int32(5)
	expected4 := []float64{}
	actual4 := SMA(prices4, timePeriod4)
	assert.Equal(t, expected4, actual4, "Test Case 4 Failed: timePeriod > len(prices) should return an empty slice")

	// Test Case 5: Edge case with timePeriod being 0
	// sma.go returns empty slice for timePeriod < 2
	prices5 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	timePeriod5 := int32(0)
	expected5 := []float64{}
	actual5 := SMA(prices5, timePeriod5)
	assert.Equal(t, expected5, actual5, "Test Case 5 Failed: timePeriod=0 should return an empty slice")

	// Test Case 6: Edge case with negative timePeriod
	// sma.go returns empty slice for timePeriod < 2
	prices6 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	timePeriod6 := int32(-2)
	expected6 := []float64{}
	actual6 := SMA(prices6, timePeriod6)
	assert.Equal(t, expected6, actual6, "Test Case 6 Failed: Negative timePeriod should return an empty slice")

	// Test Case 7: Longer series for validation
	prices7 := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	timePeriod7 := int32(4)
	// Expected: (timePeriod-1) = 3 initial 0.0 values
	// SMA[3] = (10+11+12+13)/4 = 46/4 = 11.5
	// SMA[4] = (11+12+13+14)/4 = 50/4 = 12.5
	// SMA[5] = (12+13+14+15)/4 = 54/4 = 13.5
	// SMA[6] = (13+14+15+16)/4 = 58/4 = 14.5
	// SMA[7] = (14+15+16+17)/4 = 62/4 = 15.5
	// SMA[8] = (15+16+17+18)/4 = 66/4 = 16.5
	// SMA[9] = (16+17+18+19)/4 = 70/4 = 17.5
	// SMA[10] = (17+18+19+20)/4 = 74/4 = 18.5
	expected7 := []float64{0.0, 0.0, 0.0, 11.5, 12.5, 13.5, 14.5, 15.5, 16.5, 17.5, 18.5}
	actual7 := SMA(prices7, timePeriod7)

	assert.Equal(t, len(expected7), len(actual7), "Test Case 7 Failed: Length mismatch, expected %d, got %d", len(expected7), len(actual7))
	for i := range expected7 {
		assert.InDelta(t, expected7[i], actual7[i], 0.001, "Test Case 7 Failed: Value mismatch at index %d, expected %.4f, got %.4f", i, expected7[i], actual7[i])
	}

	// Test Case 8: Values that are not monotonically increasing
	prices8 := []float64{5.0, 2.0, 8.0, 3.0, 7.0, 4.0, 9.0, 1.0, 6.0}
	timePeriod8 := int32(3)
	// Expected: (timePeriod-1) = 2 initial 0.0 values
	// SMA[2] = (5+2+8)/3 = 15/3 = 5.0
	// SMA[3] = (2+8+3)/3 = 13/3 = 4.3333...
	// SMA[4] = (8+3+7)/3 = 18/3 = 6.0
	// SMA[5] = (3+7+4)/3 = 14/3 = 4.6666...
	// SMA[6] = (7+4+9)/3 = 20/3 = 6.6666...
	// SMA[7] = (4+9+1)/3 = 14/3 = 4.6666...
	// SMA[8] = (9+1+6)/3 = 16/3 = 5.3333...
	expected8 := []float64{0.0, 0.0, 5.0, 13.0/3.0, 6.0, 14.0/3.0, 20.0/3.0, 14.0/3.0, 16.0/3.0}
	actual8 := SMA(prices8, timePeriod8)
	assert.Equal(t, len(expected8), len(actual8), "Test Case 8 Failed: Length mismatch")
	for i := range expected8 {
		assert.InDelta(t, expected8[i], actual8[i], 0.001, "Test Case 8 Failed: Value mismatch at index %d", i)
	}
}
