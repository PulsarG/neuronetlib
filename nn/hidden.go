package nn

type HiddenLayer struct {
	HiddenLayerNeurons []float64
}

func NewHiddenLayer(amountHiddenNeurons int) *HiddenLayer {
	hN := make([]float64, amountHiddenNeurons)
	return &HiddenLayer{
		HiddenLayerNeurons: hN,
	}
}
