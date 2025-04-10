# 类型转换

类型转换用于将一种数据类型的变量转换为另外一种类型的变量。

Go 语言类型转换基本格式如下：

```Go
type_name(expression)
```

type_name 为类型，expression 为表达式。

## 数值类型转换

将整型转换为浮点型：

```Go
var a int = 10
var b float64 = float64(a)
```

以下实例中将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量：

```Go
package main

import "fmt"

func main() {
   var sum int = 17
   var count int = 5
   var mean float32

   mean = float32(sum)/float32(count)
   fmt.Printf("mean 的值为: %f\n",mean)
}
```

以上实例执行输出结果为：

```text
mean 的值为: 3.400000
```

## 字符串类型转换

将一个字符串转换成另一个类型，可以使用以下语法：

```Go
var str string = "10"
var num int
num, _ = strconv.Atoi(str)
```

以上代码将字符串变量 str 转换为整型变量 num。

注意，strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 \_ 来忽略这个错误

以下实例将字符串转换为整数

```Go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("转换错误:", err)
    } else {
        fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
    }
}
```

以上实例执行输出结果为：

```text
字符串 '123' 转换为整数为：123
```

以下实例将整数转换为字符串：

```Go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := 123
    str := strconv.Itoa(num)
    fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num, str)
}
```

以上实例执行输出结果为：

```text
整数 123  转换为字符串为：'123'
```

以下实例将字符串转换为浮点数：

```Go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "3.14"
    num, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("转换错误:", err)
    } else {
        fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
    }
}
```

以上实例执行输出结果为：

```text
字符串 '3.14' 转为浮点型为：3.140000
```

以下实例将浮点数转换为字符串：

```Go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := 3.14
    str := strconv.FormatFloat(num, 'f', 2, 64)
    fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num, str)
}
```

以上实例执行输出结果为：

```text
浮点数 3.140000 转为字符串为：'3.14'
```

## 接口类型转换

接口类型转换有两种情况：类型断言和类型转换。

### 类型断言

类型断言用于将接口类型转换为指定类型，其语法为：

```Go
value.(type)
或者
value.(T)
```

其中 value 是接口类型的变量，type 或 T 是要转换成的类型。

如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功。

```Go
package main

import "fmt"

func main() {
    var i interface{} = "Hello, World"
    str, ok := i.(string)
    if ok {
        fmt.Printf("'%s' is a string\n", str)
    } else {
        fmt.Println("conversion failed")
    }
}
```

以上实例中，我们定义了一个接口类型变量 i，并将它赋值为字符串 "Hello, World"。然后，我们使用类型断言将 i 转换为字符串类型，并将转换后的值赋值给变量 str。最后，我们使用 ok 变量检查类型转换是否成功，如果成功，我们打印转换后的字符串；否则，我们打印转换失败的消息。

### 类型转换

类型转换用于将一种数据类型的变量转换为另外一种类型的变量。

```Go
T(value)
```

其中 T 为类型，value 为表达式。

在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。

```Go
package main

import "fmt"

// 定义一个接口 Writer
type Writer interface {
    Write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
    str string
}

// 实现 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
    sw.str += string(data)
    return len(data), nil
}

func main() {
    // 创建一个 StringWriter 实例并赋值给 Writer 接口变量
    var w Writer = &StringWriter{}

    // 将 Writer 接口类型转换为 StringWriter 类型
    sw := w.(*StringWriter)

    // 修改 StringWriter 的字段
    sw.str = "Hello, World"

    // 打印 StringWriter 的字段值
    fmt.Println(sw.str)
}
```

解析：

1. 定义接口和结构体：

- Writer 接口定义了 Write 方法。
- StringWriter 结构体实现了 Write 方法。

2. 类型转换：

- 将 StringWriter 实例赋值给 Writer 接口变量 w。
- 使用 w.(\*StringWriter) 将 Writer 接口类型转换为 StringWriter 类型。

3. 访问字段：

- 修改 StringWriter 的字段 str，并打印其值。

## 空接口类型

空接口 interface{} 可以持有任何类型的值。在实际应用中，空接口经常被用来处理多种类型的值。

```Go
package main

import (
    "fmt"
)

func printValue(v interface{}) {
    switch v := v.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    printValue(42)
    printValue("hello")
    printValue(3.14)
}
```

在这个例子中，printValue 函数接受一个空接口类型的参数，并使用类型断言和类型选择来处理不同的类型。

