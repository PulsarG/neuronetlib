package nn

// ReLU Leaking
func ReLU(x float64) float64 {
	if x < 0 {
		return x * 0.01
	} else {
		return x
	}
}

func ReluPrime(x float64) float64 {
	if x < 0 {
		return 0.01
	} else {
		return 1
	}
}
