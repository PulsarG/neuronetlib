package nn

import "fmt"

func TrainProcess(neuronsLayer []float64, correctNeurons []float64, Weights map[int][]float64, previusLayerNeurons []float64, LearningRate float64) (map[int][]float64, []float64) {
	e := FindErrorsInNeurons(neuronsLayer, correctNeurons)
	ep := FindErrorPreviusLayer(e, previusLayerNeurons, Weights)
	newW := SetNewWeights(previusLayerNeurons, neuronsLayer, e, Weights, LearningRate)
	newPreviusNeurons := SetNewPreviusLayer(previusLayerNeurons, ep)
	return newW, newPreviusNeurons
}

func SetNewPreviusLayer(previusLayerNeurons []float64, errorsPrevius []float64) []float64 {
	for i := range previusLayerNeurons {
		previusLayerNeurons[i] = previusLayerNeurons[i] + errorsPrevius[i]
	}
	return previusLayerNeurons
}

func SetNewWeights(previusLayerNeurons []float64, Neurons []float64, Errors []float64, Weights map[int][]float64, LearningRate float64) map[int][]float64 {
	l := len(Weights)
	NewWeights := make(map[int][]float64, l)
	for k, v := range Weights {
		NewWeights[k] = v
	}
	for i := 0; i < len(Neurons)-1; i++ {
		nw := NewWeights[i]
		for j := range nw {
			nw[j] = nw[j] + (previusLayerNeurons[j] * LearningRate * Errors[i] * ReluPrime(Neurons[i]) * previusLayerNeurons[i])
		}
		NewWeights[i] = nw
	}
	return NewWeights
}

func FindErrorPreviusLayer(Errors []float64, previusLayerNeurons []float64, Weights map[int][]float64) []float64 {
	l := len(previusLayerNeurons)
	errorsPrevius := make([]float64, l)
	for i := 0; i < l; i++ {
		for j := 0; j < len(Weights); j++ {
			x := Weights[j]
			errorsPrevius[i] += x[i] * Errors[j]
		}
	}
	return errorsPrevius
}

func FindErrorsInNeurons(neuronsLayer []float64, correctNeurons []float64) []float64 {
	l := len(neuronsLayer)
	errorsNeurons := make([]float64, l)
	for i := 0; i < l; i++ {
		errorsNeurons[i] = correctNeurons[i] - neuronsLayer[i]
	}
	x := 0.0
	for i := range errorsNeurons {
		x += errorsNeurons[i] * errorsNeurons[i]
	}
	fmt.Println("Ошибка: ", x)
	return errorsNeurons
}
