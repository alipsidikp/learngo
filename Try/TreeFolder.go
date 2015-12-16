package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type DataFolder struct {
	name  string
	items map[string]*DataFolder
}

func (d *DataFolder) gendata(list string) {
	if d.items == nil {
		d.items = make(map[string]*DataFolder)
	}

	var dir []string
	dir = strings.Split(list, "/")
	xlist := ""
	for i, xVal := range dir {
		if i > 0 {
			if xlist != "" {
				xlist = xlist + "/"
			}
			xlist = xlist + xVal
		}
	}

	if _, ok := d.items[dir[0]]; !ok {
		d.items[dir[0]] = &DataFolder{name: dir[0]}
	}

	if xlist != "" {
		// d.items[dir[0]].gendata(xlist)
	}
}

// func gendata(d *DataFolder, list string) {
// 	if d.items == nil {
// 		d.items = make(map[string]DataFolder)
// 	}

// 	var dir []string
// 	dir = strings.Split(list, "/")
// 	xlist := ""
// 	for i, xVal := range dir {
// 		if i > 0 {
// 			if xlist != "" {
// 				xlist = xlist + "/"
// 			}
// 			xlist = xlist + xVal
// 		}
// 	}

// 	if _, ok := d.items[dir[0]]; !ok {
// 		d.items[dir[0]] = DataFolder{name: dir[0]}
// 	}

// 	if xlist != "" {
// 		gendata(&(d.items[dir[0]]), xlist)
// 	}
// }

func main() {
	listpwd := []string{"contrib/file/readme.md", "etc/mysql/my.cnf", "contrib/file/config.cnf", "contrib/go.exe"}
	data := make(map[string]*DataFolder)
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
			data[dir[0]] = &DataFolder{name: dir[0]}
		}

		if xlist != "" {
			data[dir[0]].gendata(xlist)
			// gendata(data[dir[0]], xlist)
		}

	}

	fmt.Println(data)
	x, _ := json.Marshal(data)
	fmt.Println(string(x))
}
