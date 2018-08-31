package cryptopals

import b64 "encoding/base64"
import "encoding/hex"


// Task 1
func hexToBase64(hexStr string) string {
	binaryStr, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}
	return b64.StdEncoding.EncodeToString(binaryStr)
}

// Task 2
func xorHex(first string, second string) string {
	if len(first) != len(second) {
		panic("Strings must have equal length")
	}
	firstBinary, err := hex.DecodeString(first)
	if err != nil {
		panic(err)
	}
	secondBinary, err := hex.DecodeString(second)
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, len(firstBinary))
	for idx, _ := range firstBinary {
		buffer[idx] = firstBinary[idx] ^ secondBinary[idx]
	}
	return hex.EncodeToString(buffer)
}

// Task 3
func findXorChar(input string) string {
	byteInput, _ := hex.DecodeString(input)
	bestMatch := ""
	for b := 0; b < 256; b++ {
		buffer := make([]byte, len(byteInput))
		for idx, _ := range(byteInput) {
			buffer[idx] = byteInput[idx] ^ byte(b)
			// TODO: Evaluate using some metrics
		}
	}
	return bestMatch
}
