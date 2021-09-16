package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Let's print date")
	currentDate := time.Now()
	fmt.Println(currentDate)
	fmt.Println(currentDate.Format("01-02-2006 15:04:05 Monday"))
	// create a new date
	createdDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
