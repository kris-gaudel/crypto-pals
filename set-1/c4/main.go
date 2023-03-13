package main

import (
	"bytes"
	"encoding/hex"
	"strings"
	"fmt"
	"os"
	"bufio"
)

func get_score(s string) int {
	// Higher is better
	score := 0
	common := [11]string{"e", "a", "r", "i", "o", "t", "n", "s", "l", "c", " "}

	for i := range common {
		score += strings.Count(strings.ToLower(s), common[i])
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

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	best := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		best += decrypt_XOR(scanner.Text())
	}

	fmt.Println(best)
}