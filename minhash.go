package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	HASH_BASE = 5016811
)

type row []int64
type descriptor []int64
type hashFunc func(int64) int64

func buildHashFunction(a, b int64) hashFunc {
	a = a % HASH_BASE
	b = b % HASH_BASE
	return func(num int64) int64 {
		return ((a*(num%HASH_BASE))%HASH_BASE + b) % HASH_BASE
	}
}

type MinHash struct {
	hashFuncs []hashFunc
}

func NewMinHash(nFunctions uint16) MinHash {
	mh := MinHash{make([]hashFunc, nFunctions)}
	for idx := uint16(0); idx < nFunctions; idx++ {
		a := rand.Int63()
		b := rand.Int63()
		mh.hashFuncs[idx] = buildHashFunction(a, b)
	}
	return mh
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func (mh *MinHash) BuildDescriptor(items row) descriptor {
	d := make(descriptor, mh.getHashFuncsCount())
	for i := range d {
		d[i] = math.MaxInt64
	}
	for i, hashFn := range mh.hashFuncs {
		for _, item := range items {
			d[i] = min(d[i], hashFn(item))
		}
	}
	return d
}

func (mh *MinHash) Compare(first, second row) float64 {
	firstDescriptor := mh.BuildDescriptor(first)
	secondDescriptor := mh.BuildDescriptor(second)
	matchingCount := 0
	for i := range firstDescriptor {
		if firstDescriptor[i] == secondDescriptor[i] {
			matchingCount++
		}
	}
	return float64(matchingCount) / float64(len(firstDescriptor))
}

func (mh *MinHash) getHashFuncsCount() int {
	return len(mh.hashFuncs)
}

func generateSample(rowSize uint64, maxN int64) row {
	r := make(row, rowSize)
	for i := range r {
		r[i] = rand.Int63n(maxN)
	}
	return r
}

func IntersectionOverUnion(first, second row) float64 {
	firstSet := make(map[int64]bool)
	secondSet := make(map[int64]bool)
	for _, item := range first {
		firstSet[item] = true
	}
	for _, item := range second {
		secondSet[item] = true
	}
	intersection := 0
	for key := range firstSet {
		if _, ok := secondSet[key]; ok {
			intersection += 1
		}
	}
	union := len(firstSet) + len(secondSet) - intersection
	if union == 0 {
		return 0
	}
	return float64(intersection) / float64(union)
}

func main() {
	rand.Seed(42)
	nTests := 100
	var rowSize uint64 = 1000
	var maxElemInRow int64 = 1000
	var minHashSize uint16 = 20
	var err float64 = 0
	for i := 0; i < nTests; i++ {
		sample1 := generateSample(rowSize, maxElemInRow)
		sample2 := generateSample(rowSize, maxElemInRow)

		mh := NewMinHash(minHashSize)
		similiarity := mh.Compare(sample1, sample2)
		trueIoU := IntersectionOverUnion(sample1, sample2)
		fmt.Printf("MinHash: %v; IoU: %v\n", similiarity, trueIoU)
		err += math.Abs(similiarity - trueIoU)
	}
	avgError := err / float64(nTests)
	fmt.Printf("Average error: %v\n", avgError)
}
