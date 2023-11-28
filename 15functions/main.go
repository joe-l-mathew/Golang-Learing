package main

import "fmt"

func main() {
	greeter()
	result := inifniteAdder(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func greeter() {
	fmt.Println("Welcome to functions")
}

// func adder(valOne int, valTwo int) int {
// 	return valOne + valTwo
// }

func inifniteAdder(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum

}
