package operations

import "math"

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	return a / b
}

func UnaryNegative(a float64) float64 {
	return -a
}

func Root(a, root float64) float64 {
	return math.Pow(a, 1/root)
}

func SquareRoot(a float64) float64 {
	return Root(a, 2)
}
