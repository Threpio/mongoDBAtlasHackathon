package util

import "encoding/base64"

func Base64EncodeString(s string) (encoded string) {
	data := []byte(s)
	return base64.StdEncoding.EncodeToString(data)
}

func Base64DecodeString(s string) (decoded string, err error) {
	data, err := base64.StdEncoding.DecodeString(s)
	return string(data), err
}
