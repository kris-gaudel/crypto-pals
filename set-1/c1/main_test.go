package main

import (
	"testing"
)

func Test_1(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	target := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	output, err := hex_to_base64(input)
	if err != nil {
		t.Fatal(err)
	}
	if output != target {
		t.Errorf("Expected %s, got %s", target, output)
	}
}

func Test_2(t *testing.T) {
	input := "48656c6c6f20476f7068657221"
	target := "SGVsbG8gR29waGVyIQ=="
	output, err := hex_to_base64(input)
	if err != nil {
		t.Fatal(err)
	}
	if output != target {
		t.Errorf("Expected %s, got %s", target, output)
	}
}