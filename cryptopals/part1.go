package cryptopals

import b64 "encoding/base64"
import "encoding/hex"
import "math"
import "strings"


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
func xorWithByte(inputBuffer []byte, b byte) []byte {
	outputBuffer := make([]byte, len(inputBuffer))
	for idx, _ := range inputBuffer {
		outputBuffer[idx] = inputBuffer[idx] ^ b
	}
	return outputBuffer
}

func compareToEnglish(str string) float64 {
	str = strings.ToLower(str)
	divider := float64(len(str))
	if divider == 0 {
		return 0
	}
	frequencyTable := make(map[rune]float64)
	for _, ch := range str {
		frequencyTable[ch]++
	}
	for key, value := range frequencyTable {
		frequencyTable[key] = value / divider
	}
	
	var closeness float64 = 0.0
	eps := 0.0001
	for key, value := range frequencyTable {
		closeness += math.Abs(value - EnglishRuneFrequencyTable[key]) / (EnglishRuneFrequencyTable[key] + eps)
	}
	closeness /= float64(len(frequencyTable))
	return closeness
}

func findXorChar(input string) string {
	byteInput, _ := hex.DecodeString(input)
	bestString := ""
	minScore := 100500.0
	for b := 0; b < 256; b++ {
		candidate := string(xorWithByte(byteInput, byte(b)))
		currentScore := compareToEnglish(candidate)
		if currentScore < minScore {
			minScore = currentScore
			bestString = candidate
		}
	}
	return bestString
}
// End task 3
