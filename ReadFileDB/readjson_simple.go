package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type Record struct {
	EmployeeId string `json:"EmployeeId"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Email      string `json:"Email"`
}

func main() {
	fmt.Println("Decoding from variable")
	fmt.Println("======================")

	str := `{"page":1, "fruits":["apple","peach"]}`

	res := Response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.Page)
	fmt.Println(res)

	fmt.Println("======================")
	fmt.Println("Decoding from File")
	fmt.Println("======================")

	rec := []Record{}
	file, e := os.Open("E:\\data\\sample\\DataJson_Simple.json")
	if e != nil {
		fmt.Println(e.Error())
	}
	jsonParser := json.NewDecoder(file)
	e = jsonParser.Decode(&rec)

	if e != nil {
		fmt.Println("Parsing Error, ", e.Error())
	}

	fmt.Println(rec)
}
