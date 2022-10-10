package nn

type Perceptron struct {
	InputNeurons  []float64
	ResultNeurons []float64
	HiddenNeurons []HiddenLayer
	Weights       []WeightsLayer
}

func NewPerceptron(in []float64, lenResultNeurons int) *Perceptron {
	result := make([]float64, lenResultNeurons)
	return &Perceptron{
		InputNeurons:  in,
		ResultNeurons: result,
	}
}

func (p *Perceptron) SetResult(out []float64) {
	p.ResultNeurons = out
}

func (p *Perceptron) CreateHiddenNeuronsLayer(amountHiddenNeurons, amountPreviousLayer int) {
	p.Weights = append(p.Weights, *NewWeightsLayer(amountPreviousLayer, amountHiddenNeurons))
	p.HiddenNeurons = append(p.HiddenNeurons, *NewHiddenLayer(amountHiddenNeurons))
}

func (p *Perceptron) CreateLastWeights() {
	x := len(p.HiddenNeurons) - 1
	amountLastHiddenLayer := len(p.HiddenNeurons[x].HiddenLayerNeurons)
	p.Weights = append(p.Weights, *NewWeightsLayer(amountLastHiddenLayer, len(p.ResultNeurons)))
}

/* func (p *Perceptron) SetHiddenNeurons(pp *Perceptron) []float64 {
	x := len(pp.HiddenNeurons)
	l := len(pp.HiddenNeurons[0])

	nN := make([]float64, l)
	for k := 0; k < l; k++ {
		x := 0.0
		y := p.Weights[k]
		for j := range y {
			w := y[j]
			x += (previusLayerNeurons[j] * w)
		}
		nN[k] = x
	}

	y := p.HiddenNeurons[i]
	for i := range nN {
		y[i] = ReLU(y[i])
	}
	return nN
} */
