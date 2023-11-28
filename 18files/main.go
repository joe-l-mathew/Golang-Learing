package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "First rule of fight club is ******"
	file, err := os.Create("fightclub.txt")
	if err != nil {
		panic(err)
	}
	len, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Println(len)
	ReadFile("fightclub.txt")
}

func ReadFile(filePath string) {
	byteVal, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteVal))

}
