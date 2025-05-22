package indicators

import (
	"testing"

	cgotalib "github.com/blazer-org/cgo-talib"
	"github.com/stretchr/testify/assert"
)

func TestADX(t *testing.T) {
	emptyResult := []float64{}
	nilResult := []float64(nil)

	// Test Case 1: Basic valid case
	high1 := []float64{
		10.0, 10.1, 10.2, 10.0, 9.9, 10.1, 10.3, 10.4, 10.2, 10.0, // 10
		10.1, 10.2, 10.4, 10.5, 10.6, 10.5, 10.3, 10.2, 10.0, 9.8,  // 20
		10.0, 10.2, 10.3, 10.5, 10.6, 10.8, 11.0, 10.8, 10.7, 10.5, // 30
		10.6, 10.7, 10.5, 10.3, 10.1,                             // 35
	}
	low1 := []float64{
		9.8, 9.9, 10.0, 9.8, 9.7, 9.8, 10.0, 10.1, 10.0, 9.8,   // 10
		9.9, 10.0, 10.1, 10.2, 10.3, 10.2, 10.1, 10.0, 9.8, 9.6,   // 20
		9.8, 10.0, 10.1, 10.2, 10.3, 10.5, 10.6, 10.5, 10.4, 10.3, // 30
		10.4, 10.5, 10.3, 10.1, 9.9,                              // 35
	}
	close1 := []float64{
		9.9, 10.0, 10.1, 9.9, 9.8, 10.0, 10.2, 10.3, 10.1, 9.9,   // 10
		10.0, 10.1, 10.3, 10.4, 10.5, 10.4, 10.2, 10.1, 9.9, 9.7,   // 20
		9.9, 10.1, 10.2, 10.4, 10.5, 10.7, 10.9, 10.7, 10.6, 10.4, // 30
		10.5, 10.6, 10.4, 10.2, 10.0,                             // 35
	}
	timePeriod1 := int32(14) 

	expected1 := cgotalib.Adx(high1, low1, close1, timePeriod1)
	actual1 := ADX(high1, low1, close1, timePeriod1)

	assert.Equal(t, len(expected1), len(actual1), "Test Case 1 Failed: Length mismatch")
	for i := range expected1 {
		assert.InDelta(t, expected1[i], actual1[i], 0.0001, "Test Case 1 Failed: Value mismatch at index %d, expected %f, got %f", i, expected1[i], actual1[i])
	}

	// Test Case 2: Empty input slices
	actual2 := ADX([]float64{}, []float64{}, []float64{}, 14)
	assert.Equal(t, nilResult, actual2, "Test Case 2 Failed: Empty slices should return nil")

	// Test Case 3: Mismatched input slice lengths
	actual3 := ADX(high1, low1, close1[:10], 14) 
	assert.Equal(t, nilResult, actual3, "Test Case 3 Failed: Mismatched lengths should return nil")

	// Test Case 4: timePeriod < 2
	actual4 := ADX(high1, low1, close1, 1)
	assert.Equal(t, emptyResult, actual4, "Test Case 4 Failed: timePeriod < 2 should return empty slice")

	// Test Case 5: Insufficient data length for calculation
	// Required length for ADX with timePeriod=14 is 2 * 14 = 28.
	// Test with length 2*timePeriod - 1 = 27 (which should now be handled by the wrapper)
	requiredLen5 := (2*timePeriod1)
	actual5_insufficient := ADX(high1[:requiredLen5-1], low1[:requiredLen5-1], close1[:requiredLen5-1], timePeriod1)
	assert.Equal(t, emptyResult, actual5_insufficient, "Test Case 5 Failed: len < (2*timePeriod) should return empty slice")

	// Test Case 5b: Data length exactly 2*timePeriod (should now pass without panic)
	high5b := high1[:requiredLen5]
	low5b := low1[:requiredLen5]
	close5b := close1[:requiredLen5]
	expected5b := cgotalib.Adx(high5b, low5b, close5b, timePeriod1)
	actual5b := ADX(high5b, low5b, close5b, timePeriod1)
	assert.Equal(t, len(expected5b), len(actual5b), "Test Case 5b Failed: Length mismatch for exact required length")
	for i := range expected5b {
		assert.InDelta(t, expected5b[i], actual5b[i], 0.0001, "Test Case 5b Failed: Value mismatch for exact required length at index %d", i)
	}

	// Test Case 6: timePeriod = 2 (minimum valid for ADX)
	// Required length for timePeriod=2 is 2*2 = 4
	timePeriod6 := int32(2)
	high6 := []float64{10, 11, 12, 13, 14} // len 5, > 2*2=4
	low6 :=  []float64{9, 9.5, 10, 10.5, 11}
	close6:= []float64{9.5, 10, 11, 12, 13.5}
	expected6 := cgotalib.Adx(high6, low6, close6, timePeriod6)
	actual6 := ADX(high6, low6, close6, timePeriod6)
	assert.Equal(t, len(expected6), len(actual6), "Test Case 6 Failed: Length mismatch for timePeriod=2")
	for i := range expected6 {
		assert.InDelta(t, expected6[i], actual6[i], 0.0001, "Test Case 6 Failed: Value mismatch for timePeriod=2 at index %d", i)
	}

	// Test Case 7: Data that might cause DI lines to be zero (flat price action)
	// Required length for timePeriod=7 is 2*7 = 14
	high7 := []float64{10,10,10,10,10,10,10,10,10,10,10,10,10,10} // 14 points
	low7 :=  []float64{10,10,10,10,10,10,10,10,10,10,10,10,10,10}
	close7:= []float64{10,10,10,10,10,10,10,10,10,10,10,10,10,10}
	timePeriod7 := int32(7) 
	expected7 := cgotalib.Adx(high7, low7, close7, timePeriod7) 
	actual7 := ADX(high7, low7, close7, timePeriod7)
	assert.Equal(t, len(expected7), len(actual7), "Test Case 7 Failed: Length mismatch for flat data")
	for i := range expected7 {
		assert.InDelta(t, expected7[i], actual7[i], 0.0001, "Test Case 7 Failed: Value mismatch for flat data at index %d", i)
	}
}
