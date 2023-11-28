package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to rating section")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	intInput, err := strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(intInput + 1)
}
