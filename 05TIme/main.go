package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	t := time.Now()

	elapsed := t.Sub(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05")) // default format and have to use the same for date format:-  "01-02-2006 Monday 15:04:05"
	fmt.Println(elapsed)

	createdDate := time.Date(2020, time.April, 25, 23, 23, 00, 00, time.Local)

	fmt.Println("created date : ", createdDate.Format("01-02-2006 Monday"))

}
