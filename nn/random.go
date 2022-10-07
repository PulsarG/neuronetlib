package nn

import (
	"math/rand"
)

func Randweights(amountPreviousLayer int, amountNextLayer int, min float64, max float64) map[int][]float64 {
	m := make(map[int][]float64)
	for j := 0; j < amountNextLayer; j++ {
		weights := make([]float64, amountPreviousLayer)
		for i := range weights {
			weights[i] = min + rand.Float64()*(max-min)
		}
		m[j] = weights
	}
	return m
}
