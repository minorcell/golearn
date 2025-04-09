package learn

import "fmt"

func SliceTest() {
	slice := []int{1, 2, 3, 4, 5}
	printSlice(slice)

	slice = append(slice, 6)
	printSlice(slice)

	// make的三个参数：len, cap, type，分别表示长度、容量和类型
	slice2 := make([]int, len(slice), (cap(slice))*2)
	copy(slice2, slice)
	printSlice(slice2)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}