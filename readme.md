# Algoritmos

Podem ser visualizados no site [VisuAlgo](https://visualgo.net/en/sorting).

## Bubble Sort - O(n^2)

```golang
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
```

## Selection Sort - O(n^2)

```golang
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
```

## Insertion Sort - O(n^2)

```golang
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
```
