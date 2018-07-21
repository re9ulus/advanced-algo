package main

import (
	"fmt"
	"math/rand"
)

const (
	// TODO: Pass prime as param ?
	prime = 15485863
	hashFuncsRequired = 1
)


type hashFunc func(int) int


type BloomFilter struct {
	// TODO: Use real bitmask
	bitmask []bool
	checkers []hashFunc
}

func NewHashFunc(modulo int) hashFunc {
	a := rand.Int()
	b := rand.Int()
	return func(val int) int {
		// TODO: b * val - possible ovefloat
		return ((a + (b * val) % prime) % prime) % modulo
	}
}


func NewBloomFilter(size int) BloomFilter {
	// TODO: Compute required number of hash-functions
	hashFuncs := make([]hashFunc, 0)
	for i := 0; i < hashFuncsRequired; i++ {
		hashFuncs = append(hashFuncs, NewHashFunc(size))
	}
	return BloomFilter{make([]bool, size), hashFuncs}
}

func (filter BloomFilter) String() string {
	return fmt.Sprintf("%v", filter.bitmask)
}

func (filter *BloomFilter) Check(val int) bool {
	for _, checker := range filter.checkers {
		idx := checker(val)
		if !filter.bitmask[idx] {
			return false
		}
	}
	return true
}


func (filter *BloomFilter) Add(val int) {
	for _, checker := range filter.checkers {
		idx := checker(val)
		filter.bitmask[idx] = true
	}
}

func testBloomFilter() {
	filter := NewBloomFilter(10)

	fmt.Printf("default filter: %v\n", filter)
	filter.Add(5)
	fmt.Printf("updated filter: %v\n", filter)
	filter.Add(9)
	fmt.Printf("updated fitler: %v\n", filter)

	for i := 1; i < 10; i++ {
		fmt.Printf("%v is in filter: %v\n", i, filter.Check(i))
	}
}

func main() {
	fmt.Println("Start")

	testBloomFilter()

	fmt.Println("Done")
}