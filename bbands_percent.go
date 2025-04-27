package indicators

import cgotalib "github.com/blazer-org/cgo-talib"

func BbandsPercent(close []float64) []float64 {
	bbUpper, _, bbLower := cgotalib.BBands(close, int32(20), 2.0, 2.0, 0)

	var out []float64

	//  %B = (Close - Lower Band) / (Upper Band - Lower Band)
	for i := range close {
		if len(bbUpper) <= i || len(bbLower) <= i {
			out = append(out, 0)
			continue
		}
		if bbUpper[i] == bbLower[i] {
			out = append(out, 0)
		} else {
			out = append(out, (close[i]-bbLower[i])/(bbUpper[i]-bbLower[i])*100)
		}
	}
	return out
}
