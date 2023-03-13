package main

import (
	"testing"
	"os"
	"bufio"
	"strings"
)

func Test_1(t *testing.T) {
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

	contains := "Now that the party is jumping"

	if !strings.Contains(best, contains) {
		t.Errorf("Correct plaintext not found!")
	}
}