package mymath

import "math"

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Pow - функция для возведения числа в степень
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Abs - функция для вычисления абсолютного значения числа
func Abs(x float64) float64 {
	return math.Abs(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
