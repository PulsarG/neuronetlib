package nn

import (
/* "fmt" */
/* "math" */ )

// ReLU
func ReLU(x float64) float64 {
	if x < 0 {
		return 0
	} else {
		return x
	}
}

// ReLU Leaking
/* func ReLU(x float64) float64 {
	if x > 1 {
		return (1 + 0.01*(x-1))
	} else if x < 0 {
		x *= 0.01
		return x
	} else {
		return x
	}
} */

func ReluPrime(x float64) float64 {
	if x < 0 || x > 1 {
		return 0.01
	} else {
		return 1
	}
}

/* func Round(x float64, prec int) float64 {
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
} */
