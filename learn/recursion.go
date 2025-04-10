package learn

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
	}
	return result
}

func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
	if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
		return guess
	}

	newGuess := (guess + x/guess) / 2
	if newGuess == prevGuess {
		return guess
	}

	return sqrtRecursive(x, newGuess, guess, epsilon)
}

func Sqrt(x float64) float64 {
	return sqrtRecursive(x, 1.0, 0.0, 1e-9)
}

func RecursionTest() {
	fmt.Println(Factorial(5))
	fmt.Println(Sqrt(25))
}
