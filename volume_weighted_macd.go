package indicators

import (
	cgotalib "github.com/blazer-org/cgo-talib"
)

// VolumeWeightedMACD calculates the Volume Weighted MACD for a given time series.
func VolumeWeightedMACD(close []float64, volume []float64, fastPeriod, slowPeriod, signalPeriod int) ([]float64, []float64, []float64) {
	if len(close) != len(volume) {
		return nil, nil, nil // Ensure close and volume have the same length
	}

	vwClose := make([]float64, len(close))
	for i := range close {
		vwClose[i] = (close[i] * volume[i])
	}

	return cgotalib.Macd(vwClose, int32(fastPeriod), int32(slowPeriod), int32(signalPeriod))
}
