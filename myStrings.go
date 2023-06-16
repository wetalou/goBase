package goBase

import (
	"fmt"
	"strings"
)

// 测试字符串分割成数组
func TestStringSplit(s string, seq string) {
	arr := strings.Split(s, seq)
	fmt.Println("字符串分割后：", arr)
}

func TestStringContains(s string, seq string) {
	flag := strings.Contains(s, seq)
	fmt.Println("字符串是否包含：", flag)
}

func TestStringSuffixAndPrefix(s string, seq string) {
	flag := strings.HasPrefix(s, seq)
	fmt.Println("字符串是否有前缀：", flag)
	flag = strings.HasSuffix(s, seq)
	fmt.Println("字符串是否有后缀：", flag)
}

func TestStringIndex(s string, seq string) {
	flag := strings.Index(s, seq)
	fmt.Println("字符串首次出现位置：", flag)
	flag = strings.LastIndex(s, seq)
	fmt.Println("字符串最后出现的位置：", flag)
}

func TestStringJoin(arrString []string, seq string) {
	arr := strings.Join(arrString, seq)
	fmt.Println("拼接字符串后：", arr)
}

func TestStringByteAndRune(s string) {
	len := len(s)
	fmt.Println("字符串长度：", len)
	fmt.Println("下标遍历字符串：")
	for i := 0; i < len; i++ {
		fmt.Printf("(%v)%v(%c)", i, s[i], s[i])
		fmt.Print(" ")
	}

	fmt.Println()
	fmt.Println("range遍历字符串：")
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
		fmt.Print(" ")
	}
}

func TestStringChange() {
	s := "test byte"
	fmt.Println("原字符串：" + s)
	strByte := []byte(s)
	strByte[0] = 'g'
	strByte[1] = 'o'
	strByte[2] = 'o'
	strByte[3] = 'd'
	fmt.Println("转换后字符：" + string(strByte))

	s = "测试rune"
	fmt.Println("原字符串：" + s)
	strRune := []rune(s)
	strRune[0] = '验'
	strRune[1] = '证'
	fmt.Println("转换后字符：" + string(strRune))
}
