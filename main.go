package main

import (
	/* "fmt" */
	/* "math/rand" */
	"fmt"
	"neuronetbinar/nn"
	/* "time" */)

/*
	var InputNeurons []float64

var HiddenLayerNeurons []float64

var ResultNeuron []float64

var LearningRate float64
*/
var TestInput []float64
var TestResult []float64

func main() {
	TestInput = []float64{1.0, 1.0}
	TestResult = []float64{1.0, 1.0}

	NewNet := nn.NewCreater()
	NewNet.CreatePerceptron()

	NewTest := nn.NewTestData(TestInput, TestResult)
	qwer := NewTest.TestResult[0] - NewNet.GetResult()
	fmt.Println("Ошибка !^2: ", qwer*qwer)
	/* var a float64
	var b float64
	*/
	/* fmt.Println("Введите первый входной нейрон: ")
	fmt.Scanf("%f\n", &a)
	fmt.Println("Введите второй входной нейрон: ")
	fmt.Scanf("%f\n", &b)
	InputNeurons = []float64{a, b} */
	/* LearningRate = 0.1
	epoch := 1000 */

	// Создание и реализация Перцептрона

	/* NewNet := nn.NewPerceptron(InputNeurons, 2) // Созздание рамок Перцептрона - входные данные и количество выходов

	NewNet.CreateHiddenNeuronsLayer(3, 2) // Создание скрытого слоя (обязательно минимум 1 скрытый слой)
	// (принимает количество нейронов своего слоя и количество нейронов предыдущего слоя)
	// Так же создаются веса между ними
	NewNet.CreateHiddenNeuronsLayer(3, 3) // Повторить для каждого нового слоя

	NewNet.CreateLastWeights()  // Создание последних весов к выходным нейронам
	nn.SetHiddenNeurons(NewNet) // Реализация Перцептрона. Заполнение скрытых нейронов и нейронов выхода.
	*/
}
