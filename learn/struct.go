package learn

import "fmt"

type User struct {
	Name string
	Age  int
}

// 值接收者：只读，不会修改原始结构体
func (u User) SayHi() {
	fmt.Println("Hi, I'm", u.Name)
}

// 指针接收者：可以修改结构体字段
func (u *User) SetAge(age int) {
	u.Age = age
}

// 测试结构体
func StructTest() {
	u := User{Name: "Alice", Age: 25}

	u.SayHi()          // 输出：Hi, I'm Alice

	u.SetAge(30)       // 修改年龄（通过指针接收者）
	fmt.Println(u.Age) // 输出：30
}