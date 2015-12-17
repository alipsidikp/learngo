package main

import (
	"fmt"
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/json"
	"github.com/eaciit/toolkit"
)

func main() {
	dataurl := toolkit.M{}
	dataurl["Pu00231_Input.trade_date"] = "20151214"
	dataurl["Pu00231_Input.variety"] = "i"
	dataurl["Pu00231_Input.trade_type"] = "0"
	dataurl["Submit"] = "Go"
	dataurl["action"] = "Pu00231_result"

	ci := &dbox.ConnectionInfo{"E:\\data\\temp\\tempjson.json", "", "", "", nil}

	c, e := dbox.NewConnection("json", ci)
	if e != nil {
		fmt.Println(e)
	}

	e = c.Connect()
	if e != nil {
		fmt.Println(e)
	}

	e = c.NewQuery().Insert().Exec(toolkit.M{"data": dataurl})
	if e != nil {
		fmt.Println("Unable to insert: %s \n", e.Error())
	}

}
