## 数据类型：

1.  值类型：int, float, bool, string, struct, array
    1. 赋值、传参的时候，**会拷贝一份完整的数据**，彼此互不影响。
2. 引用类型：slice, map, channel, function, interface
    1. 赋值/传参时，**传的是指向底层数据的“引用”（指针）**，多个变量操作的是同一块内存。

> GO的数据类型的存储没有堆和栈之分，Go **编译器会自动决定**变量存储在栈还是堆（逃逸分析），关键点是“是否复制值”还是“复制引用”。
> 

## 数据的初始化：

在Golang中当一个变量被声明出来，但是没有显式的指定值的时候，Golang会赋予默认值：

```go
var a int
var b string
fmt.Println("a:", a, "TypeOf(a):", reflect.TypeOf(a))
fmt.Println("len(b):", len(b), "TypeOf(b):", reflect.TypeOf(b))
```

上面的代码中 a 是一个int整数类型，由于没有初始值，会被默认赋予 0 的值，而 b 则会被设置为 “” 的空字符串。因此打印的结果是：

```
a: 0 TypeOf(a): int
len(b): 0 TypeOf(b): string
```

## 常量：

```go
package main

import "fmt"

func main() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)
   println()
   println(a, b, c)   
}
```

打印结果：

```
面积为 : 50
1 false str
```

## iota:

iota，特殊常量，可以认为是一个可以被编译器修改的常量，每当 iota 在新的一行被使用时，它的值都会自动加 1；

```go
func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}

// 0 1 2 ha ha 100 100 7 8
```