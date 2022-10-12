package nn

import "fmt"

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

func (p *Perceptron) createLastWeights() {
	x := len(p.HiddenNeurons) - 1
	amountLastHiddenLayer := len(p.HiddenNeurons[x].HiddenLayerNeurons)
	p.Weights = append(p.Weights, *NewWeightsLayer(amountLastHiddenLayer, len(p.ResultNeurons)))
}

func (p *Perceptron) getInfo() {
	fmt.Println("Входы: ", p.InputNeurons)
	for i := 0; i < len(p.Weights); i++ {
		fmt.Println("Веса: ", i, p.Weights[i])
		if i < len(p.HiddenNeurons) {
			fmt.Println("Скрытый слой: ", i, p.HiddenNeurons[i])
		}
	}
	fmt.Println("Выход: ", p.ResultNeurons)
}

func (p *Perceptron) setHiddenNeurons() {
	/* x := len(p.Weights) */
	var previusLayerNeurons []float64
	/* var isLastLayer bool */
	for i := 0; i < len(p.Weights); i++ {

		var amountNextLayerNeurons int

		if i == 0 {
			previusLayerNeurons = p.InputNeurons
			amountNextLayerNeurons = len(p.HiddenNeurons[i].HiddenLayerNeurons)
		} else if i == len(p.HiddenNeurons) {
			previusLayerNeurons = p.HiddenNeurons[i-1].HiddenLayerNeurons
			/* isLastLayer = true */
			amountNextLayerNeurons = len(p.ResultNeurons)
		} else {
			previusLayerNeurons = p.HiddenNeurons[i-1].HiddenLayerNeurons
			amountNextLayerNeurons = len(p.HiddenNeurons[i].HiddenLayerNeurons)
		}

		weights := p.Weights[i].Weights
		nN := make([]float64, amountNextLayerNeurons)
		for k := 0; k < amountNextLayerNeurons; k++ {
			z := 0.0
			y := weights[k]
			for j := 0; j < len(y); j++ {
				w := y[j]
				z += (previusLayerNeurons[j] * w)
			}
			nN[k] = z
		}

		for i := 0; i < len(nN); i++ {
			fmt.Println("Скрытый нейрон до РеЛУ: ", nN[i])
			nN[i] = ReLU(nN[i])
			fmt.Println("Скрытый нейрон после РеЛУ: ", nN[i])
		}

		if i != len(p.HiddenNeurons) {
			p.HiddenNeurons[i].HiddenLayerNeurons = nN
		} else {
			p.ResultNeurons = nN
		}
		/* isLastLayer = false */
	}
}

func (p *Perceptron) LastMethodInCreatingPerceptron() {
	p.createLastWeights()
	p.setHiddenNeurons()
	p.getInfo()
}
