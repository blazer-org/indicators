package indicators

// CalculatePivot calculates the pivot points, S1, S2, R1, R2 support and resistance levels
func Pivot(high, low, close []float64) (pivot, s1, r1, s2, r2 []float64) {
	pivot = make([]float64, len(high))
	s1 = make([]float64, len(high))
	r1 = make([]float64, len(high))
	s2 = make([]float64, len(high))
	r2 = make([]float64, len(high))

	for i := 0; i < len(high); i++ {
		pivot[i] = (high[i] + low[i] + close[i]) / 3
		s1[i] = 2*pivot[i] - high[i]
		r1[i] = 2*pivot[i] - low[i]
		s2[i] = pivot[i] - (high[i] - low[i])
		r2[i] = pivot[i] + (high[i] - low[i])
	}
	return
}
