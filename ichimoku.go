package indicators

import (
	"math"
)

// IchimokuResult contains all Ichimoku lines
type IchimokuResult struct {
	TenkanSen   []float64
	KijunSen    []float64
	SenkouSpanA []float64
	SenkouSpanB []float64
	ChikouSpan  []float64
}

// highestHigh returns the highest value in a slice
func highestHigh(data []float64) float64 {
	high := data[0]
	for _, v := range data {
		if v > high {
			high = v
		}
	}
	return high
}

// lowestLow returns the lowest value in a slice
func lowestLow(data []float64) float64 {
	low := data[0]
	for _, v := range data {
		if v < low {
			low = v
		}
	}
	return low
}

// CalculateIchimoku computes Ichimoku Cloud lines from highs, lows, and closes
func Ichimoku(highs, lows, closes []float64) IchimokuResult {
	n := len(highs)
	if n != len(lows) || n != len(closes) || n < 52 {
		return IchimokuResult{}
	}

	tenkanSen := make([]float64, n)
	kijunSen := make([]float64, n)
	senkouSpanA := make([]float64, n)
	senkouSpanB := make([]float64, n)
	chikouSpan := make([]float64, n)

	// Fill with NaN initially
	for i := range tenkanSen {
		tenkanSen[i] = math.NaN()
		kijunSen[i] = math.NaN()
		chikouSpan[i] = math.NaN()
	}
	for i := range senkouSpanA {
		senkouSpanA[i] = math.NaN()
		senkouSpanB[i] = math.NaN()
	}

	for i := 0; i < n; i++ {
		// Tenkan-sen (9 periods)
		if i >= 8 {
			high := highestHigh(highs[i-8 : i+1])
			low := lowestLow(lows[i-8 : i+1])
			tenkanSen[i] = (high + low) / 2
		}

		// Kijun-sen (26 periods)
		if i >= 25 {
			high := highestHigh(highs[i-25 : i+1])
			low := lowestLow(lows[i-25 : i+1])
			kijunSen[i] = (high + low) / 2
		}

		// Senkou Span A (shifted forward 26 periods)
		if i >= 25 {
			spanA := (tenkanSen[i] + kijunSen[i]) / 2
			if i+26 < len(senkouSpanA) {
				senkouSpanA[i+26] = spanA
			}
		}

		// Senkou Span B (52 periods, shifted forward 26 periods)
		if i >= 51 {
			high := highestHigh(highs[i-51 : i+1])
			low := lowestLow(lows[i-51 : i+1])
			spanB := (high + low) / 2
			if i+26 < len(senkouSpanB) {
				senkouSpanB[i+26] = spanB
			}
		}

		// Chikou Span (close shifted backward 26 periods)
		if i >= 26 {
			chikouSpan[i-26] = closes[i]
		}
	}

	return IchimokuResult{
		TenkanSen:   tenkanSen,
		KijunSen:    kijunSen,
		SenkouSpanA: senkouSpanA,
		SenkouSpanB: senkouSpanB,
		ChikouSpan:  chikouSpan,
	}
}
