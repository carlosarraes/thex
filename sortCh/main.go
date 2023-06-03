package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SortTimes struct {
	description       string
	insertionSortTime time.Duration
	bubbleSortTime    time.Duration
	selectionSortTime time.Duration
}

func main() {
	sorts := []struct {
		description string
		data        []int
	}{
		{"100 Elementos", rand.Perm(100)},
		{"1000 Elementos", rand.Perm(1000)},
		{"10000 Elementos", rand.Perm(10000)},
		{"100000 Elementos", rand.Perm(100000)},
	}

	results := make(chan SortTimes)

	for _, sort := range sorts {
		go runSorts(sort.description, sort.data, results)
		res := <-results
		fmt.Println("Descrição:", res.description)
		fmt.Println("Insertion Sort:", res.insertionSortTime)
		fmt.Println("Selection Sort:", res.selectionSortTime)
		fmt.Println("Bubble Sort:", res.bubbleSortTime)
		fmt.Println()
	}
}

func runSorts(description string, data []int, results chan SortTimes) {
	start := time.Now()
	insertionSort(append([]int(nil), data...))
	elapsedInsertion := time.Since(start)

	start = time.Now()
	bubbleSort(append([]int(nil), data...))
	elapsedBubble := time.Since(start)

	start = time.Now()
	selectionSort(append([]int(nil), data...))
	elapsedSelection := time.Since(start)

	results <- SortTimes{description, elapsedInsertion, elapsedBubble, elapsedSelection}
}

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j = j - 1
		}
		slice[j+1] = key
	}
}

func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func selectionSort(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}
