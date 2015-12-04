package main

import (
	cr "crypto/rand"
	"encoding/base64"
	"fmt"
	mr "math/rand"
	"time"
)

func main() {
	fmt.Println("Random 01, 10 Char : ", randomstring01(10))
	fmt.Println("Random 01, 20 Char : ", randomstring01(20))

	fmt.Println("Random 02, 10 Char : ", randomstring02(10))
	fmt.Println("Random 02, 20 Char : ", randomstring02(20))
}

func randomstring01(strlen int) string {
	mr.Seed(time.Now().UTC().UnixNano())
	chars := "asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP1234567890"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[mr.Intn(len(chars))]
	}
	return string(result)
}

func randomstring02(strlen int) string {
	temp := make([]byte, strlen)
	_, e := cr.Read(temp)
	if e != nil {
		fmt.Println("error found :", e.Error())
	}
	result := base64.URLEncoding.EncodeToString(temp)
	return string(result[:strlen])
}
