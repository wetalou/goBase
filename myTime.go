package goBase

import (
	"fmt"
	"time"
)

func TestTimeNow() {
	fmt.Println("现在的时间:", time.Now())
	fmt.Println("现在的时间(unix):", time.Now().Unix())
	year, month, day := time.Now().Date()
	fmt.Println("现在的时间(date):", year, month, day)

	t := time.Now()
	year = t.Year()   // type int
	month = t.Month() // type time.Month
	day = t.Day()     // type int
	fmt.Println("现在的时间(t):", year, month, day)

	datatime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("正常日期时间(t):", datatime)
}
