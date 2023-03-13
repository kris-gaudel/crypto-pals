package main

import (
	"testing"
)

func Test_1(t *testing.T) {
	in := "1c0111001f010100061a024b53535009181c"
	against := "686974207468652062756c6c277320657965"
	target := "746865206b696420646f6e277420706c6179"
	
	output, err := fixed_XOR(in, against)
	if err != nil {
		t.Fatal(err)
	}
	if output != target {
		t.Errorf("Expected %s, got %s", target, output)
	}
}