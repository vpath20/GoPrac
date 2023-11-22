package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // -> alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //-> '-' remove password show -
	Tags     []string `json:"tags,omitempty"` //-> omitempty will dont show this if its empty
}

type details struct {
	FirstName string `json:"Fname"`
	LastName  string `json:"Lname"`
	Phone     int
	Email     string   `json:"id"`
	Password  string   `json:"-"`
	Courses   []string `json:"Languages,omitempty"`
}

func main() {
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	myCourse := []course{
		{"reactJs", 100, "youtube", "Abc123", []string{"web-dev", "json"}},
		{"mern", 200, "youtube", "Bcd234", []string{"full-stack", "front-end"}},
		{"angular", 190, "youtube", "Cde345", nil},
		{"android", 129, "youtube", "Def456", []string{"android-dev", "java", "kotlin"}},
	}
	// package the data into json format

	// finalJson, err := json.Marshal(myCourse)
	// ==== one more format===
	finalJson, err := json.MarshalIndent(myCourse, "", "\t")
	if err != nil {
		panic(err)

	}

	fmt.Printf("%s", finalJson)

	student1 := []details{
		{"Rohit", "Varshney", 123456789, "abc@gmail.com", "abc123", []string{"java", "python"}},
		{"Rohit", "Varshney", 123456789, "abc@gmail.com", "abc123", []string{"swift", "react"}},
		{"Rohit", "Varshney", 123456789, "abc@gmail.com", "abc123", nil},
	}
	data, _ := json.MarshalIndent(student1, "", "\t")
	fmt.Printf("%s", data)

}

func decodeJson() {
	// jsonDataSample := []byte(`
	// {
	// 	"Fname": "Rohit",
	// 	"Lname": "Varshney",
	// 	"Phone": 123456789,
	// 	"id": "abc@gmail.com",
	// 	"Languages": ["java","python"]
	// }
	// `)

	// var student details

	// checkValid := json.Valid(jsonDataSample)
	// if checkValid {
	// 	fmt.Println("valid JSON")
	// 	json.Unmarshal(jsonDataSample, &student)
	// 	fmt.Printf("%#v\n", student)
	// } else {
	// 	fmt.Println("Invalid JSON")
	// }

	// // some cases where we want to add data to key value

	// var mySampleData map[string]interface{}

	// json.Unmarshal(jsonDataSample, &mySampleData)

	// for k, v := range mySampleData {
	// 	fmt.Printf("key id %v and data is %v \n", k, v)
	// }

	// fmt.Printf("%#v\n", mySampleData)

	jsonDataSample2 := []byte(`
	
	{
		"coursename": "android",
		"Price": 129,
		"website": "youtube",
		"tags": [
				"android-dev",
				"java",
				"kotlin"
		]
	}
	
	`)
	var mycourse course
	checkValid2 := json.Valid(jsonDataSample2)

	if checkValid2 {
		fmt.Println("valid JSON2")
		json.Unmarshal(jsonDataSample2, &mycourse)
		fmt.Printf("%#v\n", mycourse)
	} else {
		fmt.Println("invalid JSON2")
	}

	var mySampleData2 map[string]interface{}

	json.Unmarshal(jsonDataSample2, &mySampleData2)
	for k, v := range mySampleData2 {
		fmt.Printf("key is %v value is %v\n", k, v)
	}

}
