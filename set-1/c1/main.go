package main

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
)

func hex_to_base64(hex_string string) (string, error) {
	decoded_hex, err := hex.DecodeString(hex_string)
	if err != nil {
		return "", err
	}
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(decoded_hex)))
	base64.StdEncoding.Encode(dst, decoded_hex)
	return string(dst), nil
}

func main() {
	fmt.Println(hex_to_base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	fmt.Println(hex_to_base64("1c0111001f010100061a024b53535009181c"))
}