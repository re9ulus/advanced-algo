package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

type Graph map[int]map[int]struct{}

type Matrix []int

var Empty struct{}

var wg sync.WaitGroup

var nWorkers = 4

func ReadGraph(filename string) (Graph, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	idx := 0
	nodes := make(Graph)
	maxIdx := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		idx++
		if idx == 1 {
			continue
		}
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}
		if from > maxIdx {
			maxIdx = from
		}
		if to > maxIdx {
			maxIdx = to
		}

		if _, ok := nodes[from]; !ok {
			nodes[from] = make(map[int]struct{})
		}
		if _, ok := nodes[to]; !ok {
			nodes[to] = make(map[int]struct{})
		}
		nodes[from][to] = Empty
	}
	return nodes, maxIdx
}

func BuildMatrics(nodes Graph, size int) Matrix {
	m := make(Matrix, size*size)
	for from := range nodes {
		for to := range nodes[from] {
			m[from*size+to] = 1
			m[to*size+from] = 1
		}
	}
	return m
}

func CheckVertex(m Matrix, triangles []int, size, start int) {
	defer wg.Done()
	for from := start; from < size; from += nWorkers {
		for slow := 0; slow < size; slow++ {
			row := from * size
			slowRow := slow * size
			if m[row+slow] == 1 {
				for fast := slow + 1; fast < size; fast++ {
					if m[row+fast] == 1 {
						triangles[from] += m[slowRow+fast]
					}
				}
			}
		}
	}
}

func FindTriangles(m Matrix, triangles []int, size int) {
	nWorkers := nWorkers
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go CheckVertex(m, triangles, size, i)
	}
	wg.Wait()
}

func main() {
	inputFile := "./edge.tsv"
	nodes, maxNode := ReadGraph(inputFile)
	size := maxNode + 1
	triangles := make([]int, size)
	start := time.Now()
	m := BuildMatrics(nodes, size)
	fmt.Println("Prepare: ", time.Since(start))
	start = time.Now()
	FindTriangles(m, triangles, size)
	fmt.Println("Compute: ", time.Since(start))
}
