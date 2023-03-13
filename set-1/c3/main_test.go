package main

import (
	"testing"
)

func Test_1(t *testing.T) {
	res := decrypt_XOR("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	contains := "Cooking MC's like a pound of bacon"

	if res != contains {
		t.Errorf("Correct plaintext not found!")
	}
}