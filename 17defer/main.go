package main

import "fmt"

func main() {
	defer fmt.Println("World")
	MyDefer()
	defer fmt.Println("Hello")

}

func MyDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
