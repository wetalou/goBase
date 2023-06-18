package goBase

import "fmt"

func printArr(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func TestPrintArr() {
	var arr1 [5]int
    printArr(&arr1)
    fmt.Println(arr1)

    arr2 := [...]int{2, 4, 6, 8, 10}
    printArr(&arr2)
    fmt.Println(arr2)

	arr3 := [5]int{}
	printArr(&arr3)
	fmt.Println(arr3)
}