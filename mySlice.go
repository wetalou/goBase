package goBase

import "fmt"

var arrSlice = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arrSlice[2:8]
var slice1 []int = arrSlice[0:6]             //可以简写为 var slice []int = arr[:end]
var slice2 []int = arrSlice[5:10]            //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arrSlice[0:len(arrSlice)] //var slice []int = arr[:]
var slice4 = arrSlice[:len(arrSlice)-1]      //去掉切片的最后一个元素

func TestSliceCustom() {
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arrFunc := [5]int{1, 2, 3, 4, 5}

	var s6 []int
	// 前包后不包
	s6 = arrFunc[1:4]
	fmt.Println(s6)
}

func TestSlicePrint() {
	fmt.Printf("全局变量：arr %v\n", arrSlice)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arrSlice[2:8]
	slice6 := arrSlice[0:6]              //可以简写为 slice := arr[:end]
	slice7 := arrSlice[5:10]             //可以简写为 slice := arr[start:]
	slice8 := arrSlice[0:len(arrSlice)]  //slice := arr[:]
	slice9 := arrSlice[:len(arrSlice)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)
}

func TestSliceChange() {
	data := [...]int{0, 1, 2, 3, 4, 5}

	s := data[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data)
}

func TestSliceChangeTwice() {
	data := [...]int{0, 1, 2, 3, 4, 5}

	s := data[2:4]
	s[0] += 100
	s[1] += 200

	s1 := data[2:5]
	s1[0] += 600
	s1[2] += 500

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(data)
}

func TestSliceChange3() {
	data := [...]int{0, 1, 2, 3, 4, 5}

	s := data[:]
	s1 := append(s, 7)
	
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(data)
}
