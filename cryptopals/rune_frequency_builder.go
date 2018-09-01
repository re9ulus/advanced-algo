package cryptopals

import (
	"bufio"
	"os"
)

func buildFrequencyTableFromFile(filename string) map[rune]float64 {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	table := make(map[rune]float64)
	nchars := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, char := range scanner.Text() {
			table[rune(char)]++
			nchars++
		}
	}
	for key, val := range table {
		table[key] = val / float64(nchars)
	}
	return table
}
