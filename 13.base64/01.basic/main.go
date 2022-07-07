package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	username = "admin"
	password = "admin123"
)

func main() {
	// encode
	s64 := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	fmt.Println(s64)

	// decode
	dec, _ := base64.StdEncoding.DecodeString(s64)
	pair := strings.Split(string(dec), ":")
	fmt.Println("username=" + pair[0])
	fmt.Println("password=" + pair[1])
}
