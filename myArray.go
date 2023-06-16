package goBase

import (
	"fmt"
)

// 自定义数组
func TestArrayCustom() {
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println(a, b, c, d, d[1].age)
}

// 数组循环
func TestArrayCycle() {
	arr := [...][2]int{{1, 2}, {3, 2}, {6, 5}}
	for k, v := range arr {
		fmt.Println("这是key：", k)
		for ks, vs := range v {
			fmt.Println("这是keySon：", ks, vs)
		}
	}
}
