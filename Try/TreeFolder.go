package main

import (
	// "encoding/json"
	"fmt"
	"strings"
)

type DataFolder struct {
	name  string
	items map[string]DataFolder
}

func (d *DataFolder) gendata(list string) {

}

func main() {
	listpwd := []string{"contrib/file/readme.md", "contrib/file/config.cnf", "contrib/go.exe"}
	data := make(map[string]DataFolder)
	for _, val := range listpwd {
		var dir []string
		dir = strings.Split(val, "/")
		xlist := ""
		for i, xVal := range dir {
			if i > 0 {
				if xlist != "" {
					xlist = xlist + "/"
				}
				xlist = xlist + xVal
			}
		}

		if _, ok := data[dir[0]]; !ok {
			// tempdata := DataFolder{}
			// tempdata.name = "dir[0]"
			data[dir[0]] = DataFolder{name: dir[0]}
		}

		// data[dir[0]].gendata(xlist)

	}
	// x, _ := json.Marshal(data)
	fmt.Println(data)
}
