package learn

import (
	"errors"
	"fmt"
)

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

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("负数不能开平方根")
	}
	return sqrtRecursive(x, 1.0, 0.0, 1e-9), nil
}

func RecursionTest() {
	fmt.Println(Factorial(5))
	result, err := Sqrt(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
