package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	EmployeeId string `json:"EmployeeId"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Email      string `json:"Email"`
}

func main() {

	rec := Record{}
	file, e := os.Open("E:\\data\\sample\\DataJson_Simple.json")
	if e != nil {
		fmt.Println(e.Error())
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	_, e = jsonParser.Token()
	if e != nil {
		fmt.Println(e.Error())
	}

	for jsonParser.More() {
		jsonParser.Decode(&rec)
		if e != nil {
			fmt.Println("Parsing Error, ", e.Error())
		}
		fmt.Println(rec)
	}

	_, e = jsonParser.Token()
	if e != nil {
		fmt.Println(e.Error())
	}
}
