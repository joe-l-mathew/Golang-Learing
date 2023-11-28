package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to HTTP req in golang")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	byteData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteData))
	file, err := os.Create("myHtml.html")
	if err != nil {
		panic(err)
	}
	len, err := file.Write(byteData)
	if err != nil {
		panic(err)
	}
	fmt.Println(len)
}
