package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("The enterd rating is: %v  the type is %T\n", input, input)
}
