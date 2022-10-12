package nn

import (
	"fmt"
	"math"
)

func SetHiddenNeurons(p *Perceptron) {
	x := len(p.Weights)
	var previusLayerNeurons []float64
	var isLastLayer bool
	for i := 0; i < x; i++ {
		var l int
		if i == 0 {
			previusLayerNeurons = p.InputNeurons
			l = len(p.HiddenNeurons[i].HiddenLayerNeurons)
		} else if i == len(p.HiddenNeurons) {
			previusLayerNeurons = p.HiddenNeurons[i-1].HiddenLayerNeurons
			isLastLayer = true
			l = len(p.ResultNeurons)
		} else {
			previusLayerNeurons = p.HiddenNeurons[i-1].HiddenLayerNeurons
			l = len(p.HiddenNeurons[i].HiddenLayerNeurons)
		}
		weights := p.Weights[i].Weights
		nN := make([]float64, l)
		for k := 0; k < l; k++ {
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

		if !isLastLayer {
			p.HiddenNeurons[i].HiddenLayerNeurons = nN
		} else {
			p.ResultNeurons = nN
		}
		isLastLayer = false
	}
}

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}
	return rounder / pow
}

