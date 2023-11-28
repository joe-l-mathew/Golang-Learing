package main

import "fmt"

func main() {
	var myNumber int = 23

	var ptr = &myNumber

	*ptr = *ptr + 1
	fmt.Println(myNumber)
}
