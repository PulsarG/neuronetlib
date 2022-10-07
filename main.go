package main

import (
	"fmt"
	"neuronetbinar/nn"
)

var InputNeurons []float64
var HiddenLayerNeurons []float64
var ResultNeuron []float64

func main() {
	InputNeurons = []float64{1.0, 1.0}

	SinapsWeightsFirstLayer := nn.Randweights(2, 3, 0.1, 0.5)
	SinapsWeightsSecondLayer := nn.Randweights(3, 1, 0.1, 0.5)

	fmt.Println("Веса: ")
	for key, value := range SinapsWeightsFirstLayer {
		fmt.Println("Нейрон: ", key, "Вес: ", value)
	}
	fmt.Println("Веса: ")
	for key, value := range SinapsWeightsSecondLayer {
		fmt.Println("Нейрон: ", key, "Вес: ", value)
	}

	HiddenLayerNeurons = nn.HiddenNeurons(InputNeurons, SinapsWeightsFirstLayer, 3)
	fmt.Println("Скрытые нейроны: ")
	for i := range HiddenLayerNeurons {
		fmt.Println(HiddenLayerNeurons[i])
	}
	ResultNeuron = nn.HiddenNeurons(HiddenLayerNeurons, SinapsWeightsSecondLayer, 1)

	fmt.Println("Выходной нейрон: ")
	for i := range ResultNeuron {
		fmt.Println(ResultNeuron[i])
	}
}

/* func Train() {
	ErrorResult := (ResultNeuron - TestResult) * (ResultNeuron - TestResult)
	fmt.Println("Ошибка: ", ErrorResult)
	DeltaWeights = ErrorResult * nn.ReluPrime(ResultNeuron)
	fmt.Println("Дельта: ", ErrorResult)
	for i := range SinapsWeights {
		var x = SinapsWeights[i]
		SinapsWeights[i] = x - (InputNeurons[i] * DeltaWeights * LearnRate)
		fmt.Println("Новый вес: ", SinapsWeights[i])
	}
} */
