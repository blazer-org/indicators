package indicators

// PVT calculates the Price Volume Trend (PVT) for a given time series of prices and volumes.
func PVT(prices, volumes []float64) []float64 {
	pvt := make([]float64, len(prices))

	if len(prices) != len(volumes) {
		return pvt
	}

	for i := 1; i < len(prices); i++ {
		if prices[i-1] == 0 {
			pvt[i] = pvt[i-1] // avoid division by zero
		} else {
			pvt[i] = pvt[i-1] + ((prices[i]-prices[i-1])/prices[i-1])*volumes[i]
		}
	}
	return pvt
}
