package main

import (
	/* "fmt" */
	"fmt"
	/* "math/rand" */
	/* "neuronetbinar/nn" */
	"neuronetbinar/nn"
	/* "time" */)

var InputNeurons []float64

var HiddenLayerNeurons []float64

var ResultNeuron []float64

var LearningRate float64
var TestResult []float64

func main() {
	InputNeurons = []float64{1.0, 1.0}
	TestResult = []float64{1.0, 1.0}
	/* LearningRate = 0.1
	epoch := 1000 */
	

	NewNet := nn.NewPerceptron(InputNeurons, 2)
	NewNet.CreateHiddenNeuronsLayer(3, 2)
	NewNet.CreateHiddenNeuronsLayer(3, 3)
	NewNet.CreateLastWeights()
	nn.SetHiddenNeurons(NewNet)



	fmt.Println("Входы: ", NewNet.InputNeurons)
	for i := 0; i < len(NewNet.HiddenNeurons); i++ {
		fmt.Println("Скрытый слой: ", i, NewNet.HiddenNeurons[i])
	}
	for i := 0; i < len(NewNet.Weights); i++ {
		fmt.Println("Веса: ", i, NewNet.Weights[i])
	}
	fmt.Println("Выход: ", NewNet.ResultNeurons)
}
