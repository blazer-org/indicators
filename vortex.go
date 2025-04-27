package indicators

import (
	"math"
)

// VortexResult holds the VI+ and VI- series
type VortexResult struct {
	VIPlus  []float64
	VIMinus []float64
}

// Vortex calculates the Vortex Indicator (VI+ and VI-) for given high, low, close data
func Vortex(highs, lows, closes []float64, period int) *VortexResult {
	n := len(highs)
	if n != len(lows) || n != len(closes) {
		return &VortexResult{}
	}
	if n < period+1 {
		return &VortexResult{}
	}

	vmPlus := make([]float64, n)
	vmMinus := make([]float64, n)
	tr := make([]float64, n)

	for i := 1; i < n; i++ {
		upMove := math.Abs(highs[i] - lows[i-1])
		downMove := math.Abs(lows[i] - highs[i-1])
		trueRange := math.Max(
			math.Max(highs[i]-lows[i], math.Abs(highs[i]-closes[i-1])),
			math.Abs(lows[i]-closes[i-1]),
		)

		vmPlus[i] = upMove
		vmMinus[i] = downMove
		tr[i] = trueRange
	}

	viPlus := make([]float64, n)
	viMinus := make([]float64, n)

	// Rolling sums
	for i := period; i < n; i++ {
		var sumVMPlus, sumVMMinus, sumTR float64
		for j := i - period + 1; j <= i; j++ {
			sumVMPlus += vmPlus[j]
			sumVMMinus += vmMinus[j]
			sumTR += tr[j]
		}
		viPlus[i] = sumVMPlus / sumTR
		viMinus[i] = sumVMMinus / sumTR
	}

	return &VortexResult{
		VIPlus:  viPlus,
		VIMinus: viMinus,
	}
}
