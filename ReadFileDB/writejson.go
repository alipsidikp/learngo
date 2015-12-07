package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	tempfile, _ := os.OpenFile("E:\\data\\sample\\Data_Json.temp", os.O_CREATE, 0)
	writer := json.NewEncoder(tempfile)
	io.WriteString(tempfile, "[")

	header := []string{"COLUMN01", "COLUMN02", "COLUMN03", "COLUMN04", "COLUMN05", "COLUMN06"}

	for i := 0; i < 5; i++ {
		mapRandomString := map[string]string{}
		for i, val := range header {
			mapRandomString[val] = randomstring(i + rand.Intn(15))
		}
		if i > 0 {
			io.WriteString(tempfile, ",")
		}
		writer.Encode(mapRandomString)
	}
	io.WriteString(tempfile, "]")
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
