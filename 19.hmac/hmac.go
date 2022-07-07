package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

const (
	key           = "01070709010805000600040102090804010507060506080807070708000501080701090607010506030004030307080409030701090900050808030501000506"
	CookieName    = "session-id"
	FixedPassword = "admin123"
)

func signWithSha512(password string) string {
	h := hmac.New(sha512.New, []byte(key))
	io.WriteString(h, password)
	return hex.EncodeToString(h.Sum(nil))
}

func checkSignature(signature, password string) bool {
	given, _ := hex.DecodeString(signature)
	actual, _ := hex.DecodeString(signWithSha512(password))
	return hmac.Equal(given, actual)
}
