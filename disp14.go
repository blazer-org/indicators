package indicators

import cgotalib "github.com/blazer-org/cgo-talib"

func Disp14(close []float64) []float64 {
	disp14 := make([]float64, len(close))
	sma14 := cgotalib.Sma(close, int32(14))

	for i := 0; i < len(close); i++ {
		if i < len(sma14) {
			disp14[i] = (close[i] - sma14[i]) / sma14[i] * 100
		} else {
			disp14[i] = 0
		}
	}

	return disp14
}
