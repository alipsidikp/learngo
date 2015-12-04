/*
Read JSON With Token
It Can be Powerfull if we have large JSON File and have limited memory to execute

Next will be combine to dbox library
*/
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
}

func main() {
	fmt.Println("============================================")
	fmt.Println("== With Struct  ============================")
	fmt.Println("============================================")
	rec := Record{}
	file, e := os.Open("E:\\data\\sample\\DataJson_Simple.json")
	if e != nil {
		fmt.Println(e.Error())
	}

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

	file.Close()

	fmt.Println("============================================")
	fmt.Println("== Without Struct  =========================")
	fmt.Println("== View All data even not describe before ==")
	fmt.Println("============================================")

	recmap := map[string]interface{}{}

	file, e = os.Open("E:\\data\\sample\\DataJson_Simple.json")
	if e != nil {
		fmt.Println(e.Error())
	}
	defer file.Close()

	jsonParser = json.NewDecoder(file)
	_, e = jsonParser.Token()
	if e != nil {
		fmt.Println(e.Error())
	}

	for jsonParser.More() {
		jsonParser.Decode(&recmap)
		if e != nil {
			fmt.Println("Parsing Error, ", e.Error())
		}
		fmt.Println("Read One Field : ", recmap["EmployeeId"])
		fmt.Println(recmap)
	}

	_, e = jsonParser.Token()
	if e != nil {
		fmt.Println(e.Error())
	}

}
