package main

import (
	"fmt"
	"net/url"
)

const baseUrl string = "https://lco.dev:3000/learn?coursename=docker&paymentid=dh426ddhd"

func main() {
	fmt.Println("Urls in golang")
	urlData, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The scheme name is: %v\n", urlData.Scheme)
	fmt.Printf("The host name is: %v\n", urlData.Host)
	fmt.Printf("The path name is: %v\n", urlData.Path)
	fmt.Printf("The port name is: %v\n", urlData.Port())
	fmt.Printf("The query name is: %v\n", urlData.RawQuery)

	//creating a new url

	constrctedUrl := &url.URL{
		Scheme:   "https",
		Host:     "bookmywash.com:3000",
		Path:     "/offers",
		RawQuery: "vehiclename=Audi&bodytype=sedan",
	}
	fmt.Println(constrctedUrl.String())

}
