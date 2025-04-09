# 函数

**函数声明告诉了编译器函数的名称，返回类型，和参数。**Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

```
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

- func：函数由 func 开始声明
- function_name：函数名称，参数列表和返回值类型构成了函数签名。
- parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
- return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
- 函数体：函数定义的代码集合。

函数示例：

```go
/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 声明局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}
```

## **函数返回多个值**

Go 函数可以返回多个值，例如：

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Google", "Runoob")
   fmt.Println(a, b)
}
```

## **Go 语言函数作为实参**

Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。

```go
/* 声明函数变量 */
getSquareRoot := func(x float64) float64 {
   return math.Sqrt(x)
}
/* 使用函数 */
fmt.Println(getSquareRoot(9))
```

## **Go 语言函数闭包（匿名函数)**

匿名函数是一种没有函数名的函数，通常用于在函数内部定义函数，或者作为函数参数进行传递。

以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下：

```go
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i
   }
}

func main(){
   /* nextNumber 为一个函数，函数 i 为 0 */
   nextNumber := getSequence()

   /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())

   /* 创建新的函数 nextNumber1，并查看结果 */
   nextNumber1 := getSequence()
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}
```

以下实例我们定义了多个匿名函数，并展示了如何将匿名函数赋值给变量、在函数内部使用匿名函数以及将匿名函数作为参数传递给其他函数。

```go
func main() {
    // 定义一个匿名函数并将其赋值给变量add
    add := func(a, b int) int {
        return a + b
    }

    // 调用匿名函数
    result := add(3, 5)
    fmt.Println("3 + 5 =", result)

    // 在函数内部使用匿名函数
    multiply := func(x, y int) int {
        return x * y
    }

    product := multiply(4, 6)
    fmt.Println("4 * 6 =", product)

    // 将匿名函数作为参数传递给其他函数
    calculate := func(operation func(int, int) int, x, y int) int {
        return operation(x, y)
    }

    sum := calculate(add, 2, 8)
    fmt.Println("2 + 8 =", sum)

    // 也可以直接在函数调用中定义匿名函数
    difference := calculate(func(a, b int) int {
        return a - b
    }, 10, 4)
    fmt.Println("10 - 4 =", difference)
}
```

## **Go 语言函数方法**
