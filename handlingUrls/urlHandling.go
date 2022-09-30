package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {
	//^ Parsing the URL
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	//% Query Params

	qparams := result.Query()
	fmt.Printf("%T\t", qparams)

	fmt.Println(qparams["coursename"])

	fmt.Println("")

	//& Another Method

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Deepraj",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
