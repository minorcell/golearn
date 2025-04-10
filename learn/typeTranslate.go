package learn

import (
	"fmt"
	"strconv"
)

// int -> float
func Int2Float(i int) float64 {
	return float64(i)
}

// int -> string
func Int2String(i int) string {
	return strconv.Itoa(i)
}

// float -> int
func Float2Int(f float64) int {
	return int(f)
}

// float -> string
func Float2String(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

// string -> int
func String2Int(s string) (int, error) {
	return strconv.Atoi(s)
}

// string -> float
func String2Float(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// 类型断言
func GuessType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func TypeTranslateTest() {
	// int -> float
	fmt.Println(Int2Float(1))

	// int -> string
	fmt.Println(Int2String(1))

	// float -> int
	fmt.Println(Float2Int(1.0))

	// float -> string
	fmt.Println(Float2String(1.0))

	// string -> int
	fmt.Println(String2Int("1"))

	// string -> float
	fmt.Println(String2Float("1"))

	// 类型断言
	GuessType(1)
	GuessType(1.0)
	GuessType("1")
}