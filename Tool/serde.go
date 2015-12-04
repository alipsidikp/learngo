/*
Several Type of data in Go, that we want to access in same way...
*/
package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	EmployeeId string
	FirstName  string
	LastName   string
	Age        string
	JoinDate   string
	Email      string
	Phone      string
}

type M map[string]interface{}

func main() {

	// m := M{}
	datastruct := employee{}
	datastruct.EmployeeId = fmt.Sprintf("90012019")
	datastruct.FirstName = fmt.Sprintf("Monkey")
	datastruct.LastName = fmt.Sprintf("D. Luffy")
	datastruct.Age = fmt.Sprintf("10")
	datastruct.JoinDate = fmt.Sprintf("2015-11-01")
	datastruct.Email = fmt.Sprintf("user15@yahoo.com")
	datastruct.Phone = fmt.Sprintf("085-XXX-XXX-XX")

	dmt, bdt, e := serde(datastruct)

	fmt.Println("Data Struct ===========================")
	fmt.Println(datastruct)
	fmt.Println(string(bdt))
	fmt.Println(dmt)
	fmt.Println(e)
	fmt.Println("=======================================")

	dataslice := map[string]string{"EmployeeId": "90013012", "FirstName": "AABBCC", "LastName": "DDEEFF", "Age": "10", "JoinDate": "2015-11-01", "Email": "AABB@CC.com"}

	dms, bds, e := serde(dataslice)

	fmt.Println("Data Map ===========================")
	fmt.Println(dataslice)
	fmt.Println(string(bds))
	fmt.Println(dms)
	fmt.Println(e)

	fmt.Println("=======================================")

}

func serde(v interface{}) (M, []byte, error) {
	bs, e := json.Marshal(v)
	if e != nil {
		return nil, bs, fmt.Errorf("Unable to cast to M : " + e.Error())
	}

	m := M{}

	e = json.Unmarshal(bs, &m)
	if e != nil {
		return m, bs, fmt.Errorf("Unable to uncast to M from bytes: " + e.Error())
	}

	return m, bs, nil
}
