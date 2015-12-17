package main

import (
	"encoding/json"
	"fmt"
	"github.com/eaciit/toolkit"
)

func main() {
	strJson := `[{"url":"http://www.shfe.com.cn/en/products/Gold/","nameid":"goldshfecom","calltype":"GET","sourcetype":"SourceType_Http","intervaltype":"seconds","grabinterval":"20","timeoutinterval":"5","datasetting":[{"name":"GoldTab01","rowselector":"#tab_conbox li:nth-child(1) .sjtable .listshuju tbody tr","columnsettings":[{"index":0,"alias":"Code","selector":"td:nth-child(1)"},{"index":0,"alias":"ListingDate","selector":"td:nth-child(2)"},{"index":0,"alias":"ExpirationDate","selector":"td:nth-child(3)"}],"desttype":"csv"}]}]`
	v := []toolkit.M{}
	e := json.Unmarshal([]byte(strJson), &v)

	if e != nil {
		fmt.Println("Found : ", e)
	}

	for _, valM := range v {
		fmt.Println(valM)
		b, e := json.MarshalIndent(valM, "", "  ")
		if e != nil {
			fmt.Println("Found : ", e)
		}
		fmt.Printf("\n\n\n")
		fmt.Println(string(b))
	}

}
