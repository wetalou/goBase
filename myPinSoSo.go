package goBase

import "fmt"

// 指针类型只是定义，没有初始化，指针没有具体指向的内存地址（或者说不知道指向哪个变量）
// 这个方法会报错
func TestPinNewAndMake() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}

func TestPinCreate() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}
