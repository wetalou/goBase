package goBase

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int `json:"id"`
	Gender string
	Name   string
}

func TestStructTag() {
	stu := Student{
		ID:1,
		Gender:"男",
		Name:"李青山",
	}
	fmt.Println(stu)
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json marshal 失败")
		return
	}

	fmt.Println(data)
	fmt.Printf("json str:%s\n", string(data))
}