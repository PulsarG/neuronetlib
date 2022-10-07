package nn

func HiddenNeurons(previusLayerNeurons []float64, Weights map[int][]float64, amountNextNeuronsLayer int) []float64 {
	nN := make([]float64, amountNextNeuronsLayer)
	for k := 0; k < amountNextNeuronsLayer; k++ {
		x := 0.0
		y := Weights[k]
		for j := range y {
			w := y[j]
			x += (previusLayerNeurons[j] * w)
		}
		nN[k] = x
	}
	for i := range nN {
		nN[i] = ReLU(nN[i])
	}
	return nN
}
