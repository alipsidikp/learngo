package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	sliceMapRandomString := []map[string]string{}

	header := []string{"COLUMN01", "COLUMN02", "COLUMN03", "COLUMN04", "COLUMN05", "COLUMN06"}

	for i := 0; i < 5; i++ {
		mapRandomString := map[string]string{}
		for i, val := range header {
			mapRandomString[val] = randomstring(i + rand.Intn(15))
		}
		sliceMapRandomString = append(sliceMapRandomString, mapRandomString)
	}

	fmt.Println(sliceMapRandomString)

	jRandomString, _ := json.Marshal(sliceMapRandomString)
	fmt.Println(string(jRandomString))
	_ = ioutil.WriteFile("E:\\data\\sample\\randomjson.json", jRandomString, 0)
}

func randomstring(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	chars := "asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP1234567890"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
