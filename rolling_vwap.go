package indicators

func RollingVWAP(highs, lows, closes, volumes []float64, period int) []float64 {
	n := len(highs)
	if len(lows) != n || len(closes) != n || len(volumes) != n {
		panic("all input slices must have the same length")
	}

	vwaps := make([]float64, n)
	tpvQueue := make([]float64, 0, period)
	volQueue := make([]float64, 0, period)

	var sumTPV, sumVol float64

	for i := 0; i < n; i++ {
		typicalPrice := (highs[i] + lows[i] + closes[i]) / 3
		tpv := typicalPrice * volumes[i]

		tpvQueue = append(tpvQueue, tpv)
		volQueue = append(volQueue, volumes[i])

		sumTPV += tpv
		sumVol += volumes[i]

		if len(tpvQueue) > period {
			sumTPV -= tpvQueue[0]
			sumVol -= volQueue[0]
			tpvQueue = tpvQueue[1:]
			volQueue = volQueue[1:]
		}

		if i+1 >= period {
			vwaps[i] = sumTPV / sumVol
		} else {
			vwaps[i] = 0 // or math.NaN() if you prefer
		}
	}

	return vwaps
}
