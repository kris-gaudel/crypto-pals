package main

import (
	"bytes"
	// "errors"
	"encoding/hex"
	"fmt"
	// "math"
	"strings"
)

func get_score(s string) int {
	// Higher is better
	score := 0
	goodChars := [11]string{"e", "a", "r", "i", "o", "t", "n", "s", "l", "c", " "}

	for i := range goodChars {
		score += strings.Count(strings.ToLower(s), goodChars[i])
	}

	return score
}

func apply_XOR(i rune) func(rune) rune {
	return func(char rune) rune {
		return char ^ i
	}
}

func decrypt_XOR(enc string) string{
	enc_bytes, err := hex.DecodeString(enc)
	if err != nil {
		panic(err)
	}

	high_score := 0
	var decrypt string

	for i := 0; i < 256; i++ {
		plain := string(bytes.Map(apply_XOR(rune(i)), enc_bytes))
		score := get_score(plain)
		if score > high_score {
			high_score = score
			decrypt = plain
		}
	}

	return decrypt
}

func main () {
	fmt.Println(get_score("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	fmt.Println(decrypt_XOR("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
}