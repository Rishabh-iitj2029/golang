package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Tags     []string `json:"tags,omitempty"`
	Password string   `json:"-"`
	Platform string   `json:"website"`
}

func main() {
	fmt.Println("Handling the json data")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	onlineCourses := []course{
		{"Spring Boot", 2999, []string{"backend-dev web-dev security"}, "1234", "chaiAurCode"},
		{"Go lang", 2959, []string{"backend-dev web-dev security"}, "123154", "fcc"},
		{"fastAPI", 2899, nil, "1474584", "harry"},
	}

	// final json

	finalJson, err := json.MarshalIndent(onlineCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "courseName": "Go lang",
                "Price": 2959,
                "tags": [
                        "backend-dev web-dev security"
                ],
                "website": "fcc"
        }`)
	var online course

	checkValid := json.Valid(jsonDataFromWeb)

	if(checkValid){
		fmt.Println("Valid json fomat")
		json.Unmarshal(jsonDataFromWeb, &online)
		fmt.Printf("%#v\n",online)
	}else{
		fmt.Println("JSON is not valid")
	}

	// using key value pair

	var myOnlineData map[string]interface{};
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("%v : %v\n",k, v);
	}
}
