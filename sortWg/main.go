package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	variavel := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)
	variavel4 := rand.Perm(100000)

	var wg sync.WaitGroup

	wg.Add(1)
	go runSorts("100 Elementos", variavel, &wg)

	wg.Add(1)
	go runSorts("1000 Elementos", variavel2, &wg)

	wg.Add(1)
	go runSorts("10000 Elementos", variavel3, &wg)

	wg.Add(1)
	go runSorts("100000 Elementos", variavel4, &wg)

	wg.Wait()
}

func runSorts(description string, data []int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Executando", description)

	start := time.Now()
	insertionSort(append([]int(nil), data...))
	elapsed := time.Since(start)
	fmt.Printf("[insertSort] Tempo de execução: %s\n", elapsed)

	start = time.Now()
	selectionSort(append([]int(nil), data...))
	elapsed = time.Since(start)
	fmt.Printf("[selectionSort] Tempo de execução: %s\n", elapsed)

	start = time.Now()
	bubbleSort(append([]int(nil), data...))
	elapsed = time.Since(start)
	fmt.Printf("[bubbleSort] Tempo de execução: %s\n", elapsed)
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
