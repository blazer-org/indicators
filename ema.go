package indicators

func EMA(prices []float64, span int32) []float64 {
	n := len(prices)
	if n == 0 {
		return nil
	}
	if span <= 0 {
		panic("span must be greater than 0")
	}

	alpha := 2.0 / float64(span+1) // 0.142857143
	ema := make([]float64, n)

	// Initialize EMA with first price
	ema[0] = prices[0]

	// Recursively apply EMA formula
	for i := 1; i < n; i++ {
		ema[i] = alpha*prices[i] + (1.0-alpha)*ema[i-1]
	}

	return ema
}
