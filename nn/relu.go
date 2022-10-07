package nn

func ReLU(x float64) float64 {
	if x < 0 {
		return 0
	} else {
		return x
	}
}

func ReluPrime(x float64) float64 {
	if x < 0 {
		return 0
	} else {
		return 1
	}
}