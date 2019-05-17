package main

import (
	// "bytes"
	"fmt"
	"github.com/eaciit/toolkit"
	// "net/http"
	// "net/url"
)
// test commit - using gpg
func main() {
	UrlStr := "http://www.dce.com.cn/PublicWeb/MainServlet"

	dataurl := toolkit.M{}
	dataurl["Pu00231_Input.trade_date"] = "20151214"
	dataurl["Pu00231_Input.variety"] = "i"
	dataurl["Pu00231_Input.trade_type"] = "0"
	dataurl["Submit"] = "Go"
	dataurl["action"] = "Pu00231_result"

	postdata := toolkit.M{}.Set("formvalues", dataurl)
	// dataurl := url.Values{}
	// dataurl.Add("Pu00231_Input.trade_date", "20151214")
	// dataurl.Add("Pu00231_Input.variety", "i")
	// dataurl.Add("Pu00231_Input.trade_type", "0")
	// dataurl.Add("Submit", "Go")
	// dataurl.Add("action", "Pu00231_result")

	// tdata := toolkit.Jsonify(dataurl.Encode())
	r, _ := toolkit.HttpCall(UrlStr, "POST", nil, postdata)

	fmt.Println(string(toolkit.HttpContent(r)))

	// rX, _ := http.NewRequest("POST", UrlStr, bytes.NewBuffer(tdata))

	// fmt.Println(rX)
	// client := &http.Client{}
	// resp, _ := client.Do(rX)
	// fmt.Println(string(toolkit.HttpContent(resp)))
}
