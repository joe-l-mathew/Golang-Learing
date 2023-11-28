package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web verb video")
	PerfomGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerfomGetRequest() {
	const myUrl = "http://localhost:8000/get"

	res, err := http.Get(myUrl)
	if err != nil {
		panic(res)
	}
	fmt.Println(res.StatusCode)

	defer res.Body.Close()
	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	//using strings
	var b strings.Builder
	b.Write(byteData)
	fmt.Println(b.String())
}

func PerformPostJsonRequest() {
	const myURl string = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
		{
			"name":"Joel",
			"age":22,
			"language":"golang"
		}
		`)

	res, err := http.Post(myURl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl string = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("Name", "Joel")
	data.Add("Age", "22")

	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
