package main

import (
	"fmt"
	"time"
)

func main() {

	ptime := time.Now()
	fmt.Println(ptime)

	fmt.Println(ptime.Format("01-02-2006 Monday"))

	date := time.Date(2023, time.February, 10, 12, 10, 0, 0, time.Local)
	fmt.Println(date)
	fmt.Println(date.Format("01-02-2006 Monday 15:04:05"))
}
