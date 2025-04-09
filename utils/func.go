package utils

// func can be assigned to a variable
var Add = func(a, b int) int {
	return a + b
}

func GetSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
