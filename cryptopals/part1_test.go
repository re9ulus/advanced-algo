package cryptopals

import (
	"strings"
	"testing"
	"bufio"
	"os"
	"encoding/hex"
)

import b64 "encoding/base64"

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
		t.Logf("Actual %v", actual)
		t.Errorf("Base64Xor failed")
	}
}

func TestSingleByteXor(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"
	baseFrequencyTable := buildFrequencyTableFromFile("./data/text_1.txt")
	byteInput, _ := hex.DecodeString(input)
	xorByte := findXorByte(byteInput, baseFrequencyTable)
	actual := string(xorWithByte(byteInput, xorByte))
	if actual != expected {
		t.Logf("Actual %v", actual)
		t.Errorf("TestSingleByteXor failed")
	}
}

func TestFindEncodedString(t *testing.T) {
	inputFilename := "./data/task_4.input"
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	encodedStrings := make([]string, 0)
	for scanner.Scan() {
		encodedStrings = append(encodedStrings, scanner.Text())
	}
	actual := findEncodedString(encodedStrings)
	expected := "Now that the party is jumping\n"
	if actual != expected {
		t.Logf("Actual %v", actual)
		t.Logf("Expected %v", expected)
		t.Errorf("TestFindEncodedString failed")
	}
}

func TestRepeatedXor(t *testing.T) {
	key := "ICE"
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	actual := repeatedXor(input, key)
	if expected != actual {
		t.Logf("Actual %v", actual)
		t.Errorf("TestRepeatedXor failed")
	}
}

func TestHammingDistance(t *testing.T) {
	input1, input2 := "this is a test", "wokka wokka!!!"
	expected, actual := 37, hammingDistanceStr(input1, input2)
	if expected != actual {
		t.Logf("Actual hammind distance: %v", actual)
		t.Errorf("TestHammingDistance failed")
	}
}

func TestBreakRepeatedXor(t *testing.T) {
	inputFile := "./data/task_6.input"
	file, err := os.Open(inputFile)
	checkError(err)
	scanner := bufio.NewScanner(file)
	inputData := make([]byte, 0)
	for scanner.Scan() {
		// inputData = append(inputData, )
		binaryInput, err := b64.StdEncoding.DecodeString(strings.TrimSuffix(scanner.Text(), "\n"))
		checkError(err)
		inputData = append(inputData, binaryInput...)
	}
	result := breakRepeatedXor(inputData)
	t.Logf("Result is : %v", result)
}