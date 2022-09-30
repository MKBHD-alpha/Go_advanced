package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	getRequest()

}
func getRequest() {
	const myUrl = "https://www.thunderclient.com/welcome"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	/* fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content)) */

	//^ Another way using a builder

	//STEP1: Create a Builer using the "strings" Package
	var responseString strings.Builder

	//STEP2: Storing the count of the data
	content2, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content2)

	//STEP3: Print as a String
	fmt.Println("Byte Count is :", byteCount)
	fmt.Println(responseString.String())

}
