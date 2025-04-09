package utils

import "math"

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}