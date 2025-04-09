# 生命数组

- 语法格式
```text
var arrayName [length]dataType
```
其中，arrayName 是数组的名称，size 是数组的大小，dataType 是数组中元素的数据类型。

```go
var balance [10]float32
```

1. 初始化数组
```go
var balance [10]float32
```

2. 使用列表初始化数组
```go
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

3. 使用简写方式初始化数组
```go
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

4. 不确定长度的初始化数组
```go
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

5. 访问数组元素
```go
balance[0] = 1000.0
```

6. 获取数组长度
```go
len(balance)
```

7. 遍历数组
```go
for i := 0; i < len(balance); i++ {
    fmt.Println(balance[i])
}
```

8. 使用 range 遍历数组
```go
for index, value := range balance {
    fmt.Println(index, value)
}
```

9. 综合案例
```go
func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   /* 为数组 n 初始化元素 */        
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}
```

10. 将索引为 1 和 3 的元素初始化
```go
balance3 := [5]float32{1:2.0,3:7.0}  
for k = 0; k < 5; k++ {
   fmt.Printf("balance3[%d] = %f\n", k, balance3[k] )
}
```

   