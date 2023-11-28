package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(start.Format("Monday - 02-Jan-2006: 03-3 PM"))

	createdDate := time.Date(2023, time.March, 2, 0, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)

}
