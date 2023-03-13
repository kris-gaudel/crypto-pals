package main

import (
	"errors"
	"encoding/hex"
	"fmt"
)

func XOR_bytes(buf1 []byte, buf2[]byte) ([]byte, error) {
	if len(buf1) != len(buf2) {
		return []byte{}, errors.New("Buffers must be equal length")
	}

	for i := 0; i < len(buf1_bytes); i++ {
		buf1_bytes[i] ^= buf2_bytes[i]
	}

	return buf1, nil
}

func fixed_XOR(buf1 string, buf2 string) (string, error) {
	if len(buf1) != len(buf2) {
		return "", errors.New("Buffers must be equal length")
	}

	buf1_bytes, _ := hex.DecodeString(buf1)
	buf2_bytes, _ := hex.DecodeString(buf2)

	for i := 0; i < len(buf1_bytes); i++ {
		buf1_bytes[i] ^= buf2_bytes[i]
	}

	return hex.EncodeToString(buf1_bytes), nil
}

func main() {
	res, _ := fixed_XOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	fmt.Println(res)
}