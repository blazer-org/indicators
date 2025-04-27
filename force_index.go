package indicators

// ForceIndex calculates Force Index (FI)
func ForceIndex(close, volume []float64, length int) []float64 {
	drift := 1
	if len(close) != len(volume) {
		panic("close and volume slices must have the same length")
	}

	n := len(close)
	if length <= 0 {
		length = 13
	}
	pvDiff := make([]float64, n)
	for i := range close {
		if i < drift {
			pvDiff[i] = 0.0
		} else {
			pvDiff[i] = (close[i] - close[i-drift]) * volume[i]
		}
	}

	return EMA(pvDiff[drift:], int32(length))
}
