package cryptopals

import "testing"

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expexted := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	actual := hexToBase64(input)
	if expexted != actual {
		t.Errorf("HexToBase64 failed")
	}
}

func TestXorHex(t *testing.T) {
	first := "1c0111001f010100061a024b53535009181c"
	second := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	actual := xorHex(first, second)
	if actual != expected {
		t.Errorf("Actual %v", actual)
		t.Errorf("Base64Xor failed")
	}
}
