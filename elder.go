package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
)

// ElderBull calculates the Elder Bull indicator.
// It takes a slice of closing prices and returns a slice of Elder Bull values.
//
//	Elder Bull formula:
//	  Elder Bull = EMA(Close, 13) - EMA(Close, 26)
func ElderBull(close []float64) []float64 {
	ema13 := cgotalib.Ema(close, int32(13))
	ema26 := cgotalib.Ema(close, int32(26))

	elderBull := make([]float64, len(close))

	for i := range close {
		if i < 12 {
			elderBull[i] = 0
		} else {
			elderBull[i] = ema13[i] - ema26[i]
		}
	}

	return elderBull
}

// ElderBear calculates the Elder Bear indicator.
// It takes a slice of closing prices and returns a slice of Elder Bear values.
//
//	Elder Bear formula:
//	  Elder Bear = Close - EMA(Close, 13)
func ElderBear(close []float64) []float64 {
	ema := cgotalib.Ema(close, int32(13))
	elderBear := make([]float64, len(close))

	for i, v := range close {
		if i < 12 {
			elderBear[i] = 0
		} else {
			elderBear[i] = v - ema[i]
		}
	}

	return elderBear
}
