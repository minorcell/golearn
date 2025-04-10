package learn

import (
	"fmt"
)

func ErrorTest() {
	a, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}
