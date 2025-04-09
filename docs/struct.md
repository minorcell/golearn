# 结构体

Go 语言中的结构体（struct）是一种复合数据类型，可以将多个不同类型的数据组合在一起。

## 定义结构体

```go
struct Student {
    name string
    age int
    sex string
}
```

## 使用结构体

```go
var student Student

student.name = "mcell"
student.age = 25
student.sex = "male"
```

## 初始化结构体

```go
var student Student = Student{
    name: "mcell",
    age: 25,
    sex: "male",
}
```

## 访问结构体成员

```go
fmt.Println(student.name)
fmt.Println(student.age)
fmt.Println(student.sex)
```

## 完整示例

```go
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   /* 打印 Book2 信息 */
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}

// Book 1 title : Go 语言
// Book 1 author : www.runoob.com
// Book 1 subject : Go 语言教程
// Book 1 book_id : 6495407
// Book 2 title : Python 教程
// Book 2 author : www.runoob.com
// Book 2 subject : Python 语言教程
// Book 2 book_id : 6495700
```

## 结构体作为函数参数

```go
func PrintBooks(book Books) {
    fmt.Printf("Book title : %s\n", book.title)
    fmt.Printf("Book author : %s\n", book.author)
    fmt.Printf("Book subject : %s\n", book.subject)
    fmt.Printf("Book book_id : %d\n", book.book_id)
}
```

## 🥊 Go 的结构体（`struct`） VS JavaScript 的对象（`Object`）

### 1. **定义方式**

- **Go 结构体**：
  是静态类型的复合数据结构，字段类型固定，编译时已知。

  ```go
  type User struct {
      Name string
      Age  int
  }

  user := User{Name: "Tom", Age: 30}
  ```

- **JS 对象**：
  动态、灵活，可以随时增删字段，字段名和类型都不是固定的。

  ```js
  const user = { name: "Tom", age: 30 };
  user.email = "tom@example.com"; // 可以随便加字段
  ```

---

### 2. **类型系统**

- **Go** 是强类型、静态语言。字段类型不符直接编译错误。
- **JS** 是弱类型、动态语言，字段可以随意赋值不同类型。

---

### 3. **方法绑定**

- **Go** 的方法绑定是通过 `func (u User) Method()` 实现的，可以值接收者或指针接收者。
- **JS** 对象的方法就是函数属性，可以随时增删。

  ```js
  user.sayHi = function () {
    console.log("Hi!");
  };
  ```

---

### 4. **继承 / 组合**

- **Go**：没有传统继承，推荐用组合：

  ```go
  type Person struct {
      Name string
  }

  type Employee struct {
      Person
      Salary float64
  }
  ```

- **JS**：有原型继承、类继承、混入等多种方式（有点自由过头）：

  ```js
  class Person {
    constructor(name) {
      this.name = name;
    }
  }

  class Employee extends Person {
    constructor(name, salary) {
      super(name);
      this.salary = salary;
    }
  }
  ```

---

### 5. **内存布局 / 性能**

- **Go** 更贴近底层，结构体在内存中的布局更明确，适合做高性能开发。
- **JS** 对象在引擎底层是哈希表或“形态转化对象”，性能依赖引擎优化（V8 等）。

---

### 小结

| 特性      | Go Struct        | JavaScript Object |
| --------- | ---------------- | ----------------- |
| 类型系统  | 静态、强类型     | 动态、弱类型      |
| 字段修改  | 编译期固定       | 随时可变          |
| 方法绑定  | 静态方法绑定     | 动态绑定          |
| 继承机制  | 组合             | 原型/类继承       |
| 性能/内存 | 高性能、明确布局 | 引擎优化依赖      |

## Go 中结构体绑定方法

Go 语言中，结构体可以绑定方法，方法可以是值接收者，也可以是指针接收者。

```go
package main

import (
    "fmt"
)

// 定义结构体
type User struct {
    Name string
    Age  int
}

// 绑定方法（值接收者）：只读，不会修改原始结构体
func (u User) SayHi() {
    fmt.Println("Hi, I'm", u.Name)
}

// 绑定方法（指针接收者）：可以修改结构体字段
func (u *User) SetAge(age int) {
    u.Age = age
}

func main() {
    u := User{Name: "Alice", Age: 25}

    u.SayHi()          // 输出：Hi, I'm Alice

    u.SetAge(30)       // 修改年龄（通过指针接收者）
    fmt.Println(u.Age) // 输出：30
}
```
