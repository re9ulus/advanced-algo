package cryptopals

import b64 "encoding/base64"
import (
	"encoding/hex"
	"math"
	"strings"
	"fmt"
)


// Task 1
func hexToBase64(hexStr string) string {
	binaryStr, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}
	return b64.StdEncoding.EncodeToString(binaryStr)
}
// End task 1

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
// End task 2

// Task 3
func xorWithByte(inputBuffer []byte, b byte) []byte {
	outputBuffer := make([]byte, len(inputBuffer))
	for idx, _ := range inputBuffer {
		outputBuffer[idx] = inputBuffer[idx] ^ b
	}
	return outputBuffer
}

func compareToLanguage(str string, baseFrequencyTable map[rune]float64) float64 {
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
		closeness += math.Abs(value - baseFrequencyTable[key]) / (baseFrequencyTable[key] + eps)
	}
	closeness /= float64(len(frequencyTable))
	return closeness
}

func findXorByte(byteInput []byte, frequencyTable map[rune]float64) byte {
	var bestByte byte = 0
	minScore := 100500.0
	for b := 0; b < 256; b++ {
		candidate := string(xorWithByte(byteInput, byte(b)))
		currentScore := compareToLanguage(candidate, frequencyTable)
		if currentScore < minScore {
			minScore = currentScore
			bestByte = byte(b)
		}
	}
	return bestByte
}
// End task 3

// Task 4
func findEncodedString(inputStrings []string) string {
	topString := ""
	minScore := 100500.0
	baseFrequencyTable := buildFrequencyTableFromFile("./data/text_1.txt")
	for _, str := range inputStrings {
		byteInput, _ := hex.DecodeString(str)
		xorByte := findXorByte(byteInput, baseFrequencyTable)
		encodedString := string(xorWithByte(byteInput, xorByte))
		score := compareToLanguage(encodedString, baseFrequencyTable)
		if score < minScore {
			minScore = score
			topString = encodedString
		}
	}
	return topString
}
// End task 4

// Task 5
func repeatedXor(input string, key string) string {
	byteInput, byteKey := []byte(input), []byte(key)
	buffer := make([]byte, len(byteInput))
	keyIdx := 0
	for idx, _ := range byteInput {
		buffer[idx] = byteInput[idx] ^ byteKey[keyIdx]
		keyIdx = (keyIdx + 1) % len(byteKey)
	}
	return hex.EncodeToString(buffer)
}
// End task 5

// Task 6
func computeOnesInByte(b byte) int {
	counter := 0
	checkers := []byte{1, 2, 4, 8, 16, 32, 64, 128}
	for _, chk := range checkers {
		if b & chk != 0 {
			counter++
		}
	}
	return counter
}

func hammingDistanceBinary(b1 []byte, b2 []byte) int {
	if len(b1) != len(b2) {
		panic("Input strings should have same length")
	}
	distance := 0
	for idx, _ := range b1 {
		xored := b1[idx] ^ b2[idx]
		distance += computeOnesInByte(xored)
	}
	return distance
}

func hammingDistanceStr(str1 string, str2 string) int {
	return hammingDistanceBinary([]byte(str1), []byte(str2))
}

// TODO: Pass here binary array instead of string
func findKeyLength(cyphertext []byte, minLength int, maxLength int) int {
	// TODO: Use distance btw multiple blocks
	minDistance := 5001000.0
	bestKeyLength := minLength
	for keyLength := minLength; keyLength <= maxLength; keyLength++ {
		substrLength := keyLength * 3
		substr1, substr2 := cyphertext[0:substrLength], cyphertext[substrLength:2 * substrLength]
		distance := float64(hammingDistanceBinary(substr1, substr2)) / float64(substrLength)
		if distance < minDistance {
			minDistance, bestKeyLength = distance, keyLength
		}
	}
	fmt.Println("Min distance is ", minDistance)
	return bestKeyLength
}

func splitTextByBlocks(cypher []byte, blockLength int) [][]byte {
	blocks := make([][]byte, 0)
	for idx := 0; idx < len(cypher); idx += blockLength {
		blocks = append(blocks, cypher[idx:idx + blockLength])
	}
	return blocks
}

func transpose(blocks [][]byte) [][]byte {
	transpBlocks := make([][]byte, len(blocks[0]))
	for i := range transpBlocks {
		transpBlocks[i] = make([]byte, len(blocks))
	}
	for i := range blocks {
		for j := range blocks[0] {
			transpBlocks[j][i] = blocks[i][j]
		}
	}
	return transpBlocks
}

func breakBlocks(blocks [][]byte, frequencyTable map[rune]float64) [][]byte { 
	decodedBlocks := make([][]byte, len(blocks))
	for i, block := range blocks {
		xorByte := findXorByte(block, frequencyTable)
		decodedBlocks[i] = xorWithByte(block, xorByte)
	}
	return decodedBlocks
}

func mergeBlocks(blocks [][]byte) []byte {
	// TODO: allocate memory
	merged := make([]byte, 0)
	for i := range blocks {
		for j := range blocks[0] {
			merged = append(merged, blocks[i][j])
		}
	}
	return merged
}

func breakRepeatedXor(cyphertext []byte) string {
	// 1. Get size of block
	// b64.StdEncoding.DecodeString(cypher[0:keyLength])
	blockSize := findKeyLength(cyphertext, 2, 40)
	fmt.Printf("Expected block size is %v\n", blockSize)
	// 2. Split by blocks and transpose
	blocks := splitTextByBlocks(cyphertext, blockSize)
	// 3. Break blocks
	transposedBlocks := transpose(blocks)
	frequencyTable := buildFrequencyTableFromFile("./data/text_1.txt")
	decodedBlocks := breakBlocks(transposedBlocks, frequencyTable)
	// 4. Transpose back
	blocks = transpose(decodedBlocks)
	// 5. Return string
	return string(mergeBlocks(blocks))
} 
// End task 6
