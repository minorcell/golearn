package main

import (
	"fmt"

	"golearn/ch1"
	"golearn/learn"
	"golearn/utils"
	"reflect"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(utils.GetNow())
	ch1.PrintArgs()
	ch1.PrintArgs2()

	// 创建学习信息 打印学生信息
	var student = map[string]string{
		"name": "mcell",
		"age":  "25",
		"sex":  "male",
	}

	for k, v := range student {
		fmt.Println(k, v)
	}

	score := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range score {
		sum += v
	}
	fmt.Println("总成绩:", sum)

	var a int
	var b string
	fmt.Println("a:", a, "TypeOf(a):", reflect.TypeOf(a))
	fmt.Println("len(b):", len(b), "TypeOf(b):", reflect.TypeOf(b))

	var i, j int
	var primes []int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			primes = append(primes, i)
		}
	}

	fmt.Println("primes:", primes)

	fmt.Println("Max:", utils.Max(1, 2))
	fmt.Println("Min:", utils.Min(1, 2))
	fmt.Println("Abs:", utils.Abs(-1))
	fmt.Println("Pow:", utils.Pow(2, 3))

	// func can be assigned to a variable
	fmt.Println("Add:", utils.Add(1, 2))

	// Get the sequence generator function and then call it
	sequenceGenerator := utils.GetSequence()
	fmt.Println("GetSequence first call:", sequenceGenerator())
	fmt.Println("GetSequence second call:", sequenceGenerator())
	fmt.Println("GetSequence third call:", sequenceGenerator())

	// array
	blances := [...]int{1: 0, 2: 1, 3: 2, 4: 3, 5: 4}

	for index, value := range blances {
		fmt.Printf("index:%d, value:%d\n", index, value)
	}

	learn.StructTest()
	learn.SliceTest()
	learn.MapTest()
	learn.RecursionTest()
	learn.TypeTranslateTest()
	learn.InterfaceTest1()
	learn.InterfaceTest2()
	learn.ErrorTest()
	learn.ConcurrentTest()
}

// 单行注释

/*
多行注释
*/

/*
1.go run ,go run
2.go build ,./,go build
*/
