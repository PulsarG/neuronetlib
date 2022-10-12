package nn

import "fmt"

type Creater struct {
	InputNeurons       int
	OutputNeurons      int
	AmountHiddenLayers int
	Net                Perceptron
}

func NewCreater() *Creater {
	var a, b, c int
	fmt.Println("Введите количество входных нейронов: ")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Введите количество выходных нейронов: ")
	fmt.Scanf("%d\n", &b)
	fmt.Println("Введите количество скрытых слоев: ")
	fmt.Scanf("%d\n", &c)
	return &Creater{
		InputNeurons:       a,
		OutputNeurons:      b,
		AmountHiddenLayers: c,
	}
}

func (c *Creater) CreatePerceptron() {

	InputNeurons := make([]float64, c.InputNeurons)
	for i := 0; i < c.InputNeurons; i++ {
		var a float64
		fmt.Println("Введите входной нейрон", i+1)
		fmt.Scanf("%f\n", &a)
		InputNeurons[i] = a
	}

	NewNet := NewPerceptron(InputNeurons, c.OutputNeurons)
	if c.AmountHiddenLayers != 0 {
		for i := 0; i < c.AmountHiddenLayers; i++ {
			var x int
			fmt.Println("Введите количество нейронов скрытого слоя номер", i+1)
			fmt.Scanf("%d\n", &x)
			if i == 0 {
				NewNet.CreateHiddenNeuronsLayer(x, c.InputNeurons)
			} else {
				NewNet.CreateHiddenNeuronsLayer(x, len(NewNet.HiddenNeurons[i-1].HiddenLayerNeurons))
			}
		}
	}
	/* NewNet.CreateLastWeights()
	SetHiddenNeurons(NewNet)
	NewNet.SetHiddenNeurons()
	NewNet.GetInfo() */
	NewNet.LastMethodInCreatingPerceptron()

	c.Net = *NewNet
}

func (c *Creater) GetResult() float64 {
	return c.Net.ResultNeurons[0]
}
