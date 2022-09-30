package main

import (
	"encoding/json"
	"fmt"
)

// Creating Alias for better API readability using “
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //& Hides the field (doesn't reflect in api calls)
	Tags     []string `json:"tags,omitempty"` //^ if empty don't throw the field
}

func main() {
	fmt.Println("Welcome")
	//EncodeJson()
	DecodeJson()

}

// EncodeJson FUnction to Encode Json
func EncodeJson() {
	myCourses := []course{
		{"Flutter Bootcamp", 299, "deeprajbaidya.in", "baidya", []string{"App Development", "Flutter"}},
		{"MongoDB Bootcamp", 299, "deeprajbaidya.in", "baidya", []string{"App Development", "mongodb"}},
		{"Android Bootcamp", 299, "deeprajbaidya.in", "baidya", nil},
	}
	/* finalJson,err := json.Marshal(myCourses)  //@ Have to pass the Interface(instance of the struct) */
	finalJson, err := json.MarshalIndent(myCourses, "", "\t") //% For better readability
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

// Comaparing Json data

func DecodeJson() {
	dataFromTheWeb := []byte(`  
	{
		"coursename": "Flutter Bootcamp",
		"Price": 299,
		"website": "deeprajbaidya.in",
		"tags": ["App Development","Flutter"]
	}
`)
	var theCourses course                    //@ Creating the instance to store values
	checkValid := json.Valid(dataFromTheWeb) //% To check if the json Data is Valid wrt what weve rpovided.
	if checkValid {
		fmt.Println("JSON was Valid")
		/* Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
		Unmarshal uses the inverse of the encodings that Marshal uses, allocating maps, slices, and pointers as necessary, with the following additional rules: */
		json.Unmarshal(dataFromTheWeb, &theCourses)
		fmt.Printf("%#v\n", theCourses) //? inorder to print the INTERFACES

	} else {
		fmt.Println("JSON was not Valid")
	}

}
