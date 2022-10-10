package nn

type WeightsLayer struct {
	Weights map[int][]float64
}

func NewWeightsLayer(amountPreviousLayer, amountNextLayer int) *WeightsLayer {
	nw := Randweights(amountPreviousLayer, amountNextLayer, -0.5, 0.5)

	return &WeightsLayer{
		Weights: nw,
	}
}
