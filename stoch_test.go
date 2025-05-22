package indicators

import (
	"testing"

	cgotalib "github.com/blazer-org/cgo-talib" // For direct comparison
	"github.com/stretchr/testify/assert"
)

func TestStochasticOscillator(t *testing.T) {
	emptyResult := StochasticResult{StochK: []float64{}, StochKSignal: []float64{}}
	const SmaMatype int32 = 0 // TA-Lib MA_Type for SMA is 0

	// Test Case 1: Basic valid case
	high1 := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	low1 :=  []float64{1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11}
	close1:= []float64{5,  5.5,6,  6.5,7,  7.5,8,  8.5,9,  9.5,10}
	window1 := 5        // fastKPeriod
	smoothWindow1 := 3  // fastDPeriod

	// Use cgo-talib directly to get the expected output for comparison.
	// Corrected to use Stochf (lowercase f)
	expectedK1, expectedD1 := cgotalib.Stochf(high1, low1, close1, int32(window1), int32(smoothWindow1), SmaMatype)
	actual1 := StochasticOscillator(high1, low1, close1, window1, smoothWindow1)

	assert.Equal(t, len(expectedK1), len(actual1.StochK), "Test Case 1 StochK Failed: Length mismatch")
	assert.Equal(t, len(expectedD1), len(actual1.StochKSignal), "Test Case 1 StochKSignal Failed: Length mismatch")

	for i := range expectedK1 {
		assert.InDelta(t, expectedK1[i], actual1.StochK[i], 0.0001, "Test Case 1 StochK Failed: Value mismatch at index %d, expected %f, got %f", i, expectedK1[i], actual1.StochK[i])
	}
	for i := range expectedD1 {
		assert.InDelta(t, expectedD1[i], actual1.StochKSignal[i], 0.0001, "Test Case 1 StochKSignal Failed: Value mismatch at index %d, expected %f, got %f", i, expectedD1[i], actual1.StochKSignal[i])
	}

	// Test Case 2: Empty input slices
	actual2 := StochasticOscillator([]float64{}, []float64{}, []float64{}, 5, 3)
	assert.Equal(t, emptyResult, actual2, "Test Case 2 Failed: Empty slices not handled correctly")

	// Test Case 3: window < 1
	actual3 := StochasticOscillator(high1, low1, close1, 0, 3)
	assert.Equal(t, emptyResult, actual3, "Test Case 3 Failed: window < 1 not handled")

	// Test Case 4: smoothWindow < 1
	actual4 := StochasticOscillator(high1, low1, close1, 5, 0)
	assert.Equal(t, emptyResult, actual4, "Test Case 4 Failed: smoothWindow < 1 not handled")

	// Test Case 5: Insufficient data length for calculation
	// Required length for window=5, smoothWindow=3 is (5-1) + (3-1) + 1 = 4 + 2 + 1 = 7.
	shortHigh :=  []float64{10,11,12,13,14,15}    // Length 6
	shortLow :=   []float64{1, 2, 3, 4, 5, 6}     // Length 6
	shortClose := []float64{5,5.5,6,6.5,7,7.5}   // Length 6
	actual5 := StochasticOscillator(shortHigh, shortLow, shortClose, 5, 3)
	assert.Equal(t, emptyResult, actual5, "Test Case 5 Failed: Insufficient data length (len < window+smoothWindow-1) not handled")
	
	// Test Case 6: Data length just enough for calculation
	// Required length = window + smoothWindow - 1. For (5,3) this is 7.
	high6 :=  []float64{10,11,12,13,14,15,16}    // Length 7
	low6  :=  []float64{1,2,3,4,5,6,7}          // Length 7
	close6:=  []float64{5,5.5,6,6.5,7,7.5,8}    // Length 7
	window6 := 5
	smoothWindow6 := 3

	expectedK6, expectedD6 := cgotalib.Stochf(high6, low6, close6, int32(window6), int32(smoothWindow6), SmaMatype)
	actual6 := StochasticOscillator(high6, low6, close6, window6, smoothWindow6)

	assert.Equal(t, len(expectedK6), len(actual6.StochK), "Test Case 6 StochK Failed: Length mismatch")
	assert.Equal(t, len(expectedD6), len(actual6.StochKSignal), "Test Case 6 StochKSignal Failed: Length mismatch")

	if len(expectedK6) > 0 { 
		for i := range expectedK6 {
			assert.InDelta(t, expectedK6[i], actual6.StochK[i], 0.0001, "Test Case 6 StochK Failed: Value mismatch at index %d", i)
		}
	}
	if len(expectedD6) > 0 { 
		for i := range expectedD6 {
			assert.InDelta(t, expectedD6[i], actual6.StochKSignal[i], 0.0001, "Test Case 6 StochKSignal Failed: Value mismatch at index %d", i)
		}
	}

	// Test Case 7: Mismatched input slice lengths (close shorter)
	mismatchHigh7  := []float64{10,11,12,13,14}
	mismatchLow7   := []float64{1,2,3,4,5}
	mismatchClose7 := []float64{5,5.5,6} // Shorter
	actual7 := StochasticOscillator(mismatchHigh7, mismatchLow7, mismatchClose7, 3, 2)
	assert.Equal(t, emptyResult, actual7, "Test Case 7 Failed: Mismatched input slice lengths (close shorter) not handled")

	// Test Case 8: Mismatched input slice lengths (high shorter)
	mismatchHigh8  := []float64{10,11}
	mismatchLow8   := []float64{1,2,3}
	mismatchClose8 := []float64{5,5.5,6}
	actual8 := StochasticOscillator(mismatchHigh8, mismatchLow8, mismatchClose8, 2,1) 
	assert.Equal(t, emptyResult, actual8, "Test Case 8 Failed: Mismatched input slice lengths (high shorter) not handled")

	// Test Case 9: Window = 1, SmoothWindow = 1 (minimum valid periods for StochF)
	// Required length = 1 + 1 - 1 = 1
	high9   := []float64{10, 11, 12}
	low9    := []float64{1, 2, 3}
	close9  := []float64{5, 6, 7}
	window9 := 1
	smoothWindow9 := 1
	expectedK9, expectedD9 := cgotalib.Stochf(high9, low9, close9, int32(window9), int32(smoothWindow9), SmaMatype)
	actual9 := StochasticOscillator(high9, low9, close9, window9, smoothWindow9)
	assert.Equal(t, len(expectedK9), len(actual9.StochK), "Test Case 9 StochK Failed: Length mismatch")
	assert.Equal(t, len(expectedD9), len(actual9.StochKSignal), "Test Case 9 StochKSignal Failed: Length mismatch")
	if len(expectedK9) > 0 {
		for i := range expectedK9 {
			assert.InDelta(t, expectedK9[i], actual9.StochK[i], 0.0001, "Test Case 9 StochK Failed: Value mismatch at index %d", i)
		}
	}
	if len(expectedD9) > 0 { 
		for i := range expectedD9 {
			assert.InDelta(t, expectedD9[i], actual9.StochKSignal[i], 0.0001, "Test Case 9 StochKSignal Failed: Value mismatch at index %d", i)
		}
	}
	
	// Test Case 10: All input values are the same (potential division by zero if not handled by TA-Lib)
	high10 := []float64{10, 10, 10, 10, 10, 10, 10}
	low10  := []float64{10, 10, 10, 10, 10, 10, 10}
	close10:= []float64{10, 10, 10, 10, 10, 10, 10}
	window10 := 5
	smoothWindow10 := 3
	expectedK10, expectedD10 := cgotalib.Stochf(high10, low10, close10, int32(window10), int32(smoothWindow10), SmaMatype)
	actual10 := StochasticOscillator(high10, low10, close10, window10, smoothWindow10)
	assert.Equal(t, len(expectedK10), len(actual10.StochK), "Test Case 10 StochK Failed: Length mismatch")
	assert.Equal(t, len(expectedD10), len(actual10.StochKSignal), "Test Case 10 StochKSignal Failed: Length mismatch")
	if len(expectedK10) > 0 {
		for i := range expectedK10 {
			assert.InDelta(t, expectedK10[i], actual10.StochK[i], 0.0001, "Test Case 10 StochK Failed: Value mismatch at index %d", i)
		}
	}
	if len(expectedD10) > 0 {
		for i := range expectedD10 {
			assert.InDelta(t, expectedD10[i], actual10.StochKSignal[i], 0.0001, "Test Case 10 StochKSignal Failed: Value mismatch at index %d", i)
		}
	}
}
