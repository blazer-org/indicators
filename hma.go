package indicators

import (
	"math"

	cgotalib "github.com/blazer-org/cgo-talib"
)

// HMA calculates the Hull Moving Average for the given data and period.
func HMA(data []float64, length int) []float64 {
	if length < 1 {
		length = 10
	}

	half := length / 2
	sqrtLength := int(math.Sqrt(float64(length)))

	wmaFull := cgotalib.Wma(data, int32(length))
	wmaHalf := cgotalib.Wma(data, int32(half))

	diff := Subtract(Multiply(wmaHalf, 2), wmaFull)
	hma := cgotalib.Wma(diff, int32(sqrtLength))

	return hma
}

// Subtract two series element-wise
func Subtract(a, b []float64) []float64 {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	result := make([]float64, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] - b[i]
	}
	return result
}

// Multiply a series by a scalar
func Multiply(series []float64, scalar float64) []float64 {
	result := make([]float64, len(series))
	for i := range series {
		result[i] = series[i] * scalar
	}
	return result
}
