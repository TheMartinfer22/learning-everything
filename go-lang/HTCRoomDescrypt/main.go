package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
)

func main() {
	ask3, _ := base32.StdEncoding.DecodeString("MJQXGZJTGIQGS4ZAON2XAZLSEBRW63LNN5XCA2LOEBBVIRRHOM======")
	println(string(ask3))

	ask4, _ := base64.StdEncoding.DecodeString("RWFjaCBCYXNlNjQgZGlnaXQgcmVwcmVzZW50cyBleGFjdGx5IDYgYml0cyBvZiBkYXRhLg==")
	println(string(ask4))

	ask5, _ := hex.DecodeString("68657861646563696d616c206f72206261736531363f")
	println(string(ask5))

	ask6 := decodeROT13("Ebgngr zr 13 cynprf!")
	println(string(ask6))
}

func decodeROT13(s string) string {
	var decoded string
	for _, char := range s {
		switch {
		case 'a' <= char && char <= 'z':
			decoded += string((char-'a'+13)%26 + 'a')
		case 'A' <= char && char <= 'Z':
			decoded += string((char-'A'+13)%26 + 'A')
		default:
			decoded += string(char)
		}
	}
	return decoded
}
