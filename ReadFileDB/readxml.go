/*
Read XML With Token
It Can be Powerfull if we have large XML File and have limited memory to execute

Next will be combine to dbox library
*/

package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type record struct {
	EmployeeId string `xml:"EmployeeId"`
	FirstName  string `xml:"FirstName"`
	LastName   string `xml:"LastName"`
	Age        string `xml:"Age"`
	JoinDate   string `xml:"JoinDate"`
	Email      string `xml:"Email"`
	Phone      string `xml:"Phone"`
}

type record01 struct {
	value interface{}
}

func main() {

	var e error
	var rec record

	file, e := os.OpenFile("E:\\data\\sample\\Data_F1.xml", os.O_APPEND|os.O_CREATE, 0666)
	if e != nil {
		fmt.Println(e.Error())
	}

	reader := xml.NewDecoder(file)
	i := 1
	for {

		t, _ := reader.Token()
		if t == nil {
			break
		}

		switch el := t.(type) {
		case xml.StartElement:

			if el.Name.Local == "record" {
				reader.DecodeElement(&rec, &el)
				fmt.Println(rec)
			}

		default:
		}

		i += 1
		if i > 10 {
			break
		}
	}
}
