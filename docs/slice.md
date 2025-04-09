# 切片

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

## 定义切片

你可以声明一个未指定大小的数组来定义切片：

```go
var identifier []type
```

切片不需要说明长度。

或使用 make() 函数来创建切片:

```go
var slice1 []type = make([]type, len)

// 简写
slice1 := make([]type, len)
```

也可以指定容量，其中 capacity 为可选参数。

```go
make([]T, length, capacity)
```

这里 len 是数组的长度并且也是切片的初始长度。

## 切片初始化

```go
s :=[] int {1,2,3 }
```

直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

```go
s := arr[:]
```

初始化切片 s，是数组 arr 的引用。

```go
s := arr[startIndex:endIndex]
```

将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。

```go
s := arr[startIndex:]
```

默认 endIndex 时将表示一直到 arr 的最后一个元素。

```go
s := arr[:endIndex]
```

默认 startIndex 时将表示从 arr 的第一个元素开始。

## len() 和 cap() 函数

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

```go
package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)

   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
// len=3 cap=5 slice=[0 0 0]
```

## 空(nil)切片

一个切片在未初始化之前默认为 nil，长度为 0，实例如下：

```go
package main

import "fmt"

func main() {
   var numbers []int

   printSlice(numbers)

   if(numbers == nil){
      fmt.Printf("切片是空的")
   }
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

输出：

```text
len=0 cap=0 slice=[]
切片是空的
```

## 切片截取

可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，实例如下：

```go
package main

import "fmt"

func main() {
   /* 创建切片 */
   numbers := []int{0,1,2,3,4,5,6,7,8}
   printSlice(numbers)

   /* 打印原始切片 */
   fmt.Println("numbers ==", numbers)

   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])

   /* 默认下限为 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])

   /* 默认上限为 len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers1 := make([]int,0,5)
   printSlice(numbers1)

   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
   number2 := numbers[:2]
   printSlice(number2)

   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   number3 := numbers[2:5]
   printSlice(number3)

}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

## append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法

```go
package main

import "fmt"

func main() {
   var numbers []int
   printSlice(numbers)

   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers)

   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers)

   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

输出：

```text
len=0 cap=0 slice=[]
len=1 cap=1 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=6 slice=[0 1 2 3 4]
len=5 cap=12 slice=[0 1 2 3 4]
```

## Slice 和 Array 的区别

这个问题是 Go 学习路上的**“送命题”之一**——很多初学者刚开始会把 `slice` 和 `array` 搞混，结果调试时发现明明改了值却没变，或者复制时出了鬼。

---

## 🥷 Go 中 Slice 和 Array 的区别

| 特性         | 数组（Array）                | 切片（Slice）               |
| ------------ | ---------------------------- | --------------------------- |
| 长度         | 固定，编译时已知             | 动态，可增长缩小            |
| 内存分配     | 值类型，栈或静态内存         | 引用类型，指向底层数组      |
| 声明方式     | `[3]int{1, 2, 3}`            | `[]int{1, 2, 3}`            |
| 拷贝行为     | 拷贝整个数组（深拷贝）       | 拷贝引用（浅拷贝）          |
| 函数传参     | 拷贝整个数组                 | 引用传递，修改会影响原切片  |
| 可否追加元素 | ❌ 不可                      | ✅ 可用 `append()` 动态扩展 |
| 常用吗       | 不常用（除非你非常在意性能） | ✅ 是 Go 中主流容器         |

---

## 👀 举例比较：Array vs Slice

### ✅ Array（数组）

```go
arr := [3]int{1, 2, 3}
fmt.Println(arr[0]) // 输出 1

arr2 := arr
arr2[0] = 99
fmt.Println(arr[0]) // 还是 1，说明是深拷贝
```

### ✅ Slice（切片）

```go
s := []int{1, 2, 3}
fmt.Println(s[0]) // 输出 1

s2 := s
s2[0] = 99
fmt.Println(s[0]) // 输出 99，共用底层数组
```

---

## 🔬 底层原理：Slice 是“数组视图”

切片底层是个结构体，包含三部分：

```go
type slice struct {
    ptr *T   // 指向底层数组的指针
    len int  // 当前长度
    cap int  // 容量（底层数组的长度）
}
```

所以你对切片的修改会**直接影响底层数组**，而多个切片可以共享这个数组（危险，慎用 ⚠️）。

---

## 🔧 append 示例（切片扩容）

```go
nums := []int{1, 2, 3}
nums = append(nums, 4, 5) // 扩容
fmt.Println(nums)         // [1 2 3 4 5]
```

⚠️ 注意：`append` 可能触发**底层数组复制和扩容**，这时候切片就不会再引用原数组了。

---

## 🎯 总结口诀（记不住看这）

> 数组固定没法变，拷贝复制不共源；  
> 切片灵活靠底层，改个值全体变脸。  
> append 爽但要警惕，扩容后你俩不是“亲”。
