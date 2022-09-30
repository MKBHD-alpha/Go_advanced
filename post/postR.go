package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//performPostJsonRequest()
	performPostFormRequest()
}
func performPostJsonRequest() {
	const myURL = "http://localhost:1111/post"

	//* Fake JSON Payload
	requestBody := strings.NewReader(`
	{
		"coursename":"Go With Golang",
		"price":0,
		"platform":"learncodeonline.in"
	}
	`)
	response, err := http.Post(myURL, "application/json", requestBody) //@ Destination : Type :requestbody
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

/*


Performing POST Request for FORMS

*/

func performPostFormRequest() {
	const myURL = "http://localhost:1111/postform"
	//% Creating fake formdata
	data := url.Values{}
	data.Add("firstname", "Deepraj")
	data.Add("lastname", "Baidya")
	data.Add("age", "19")

	response, err := http.PostForm(myURL, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
