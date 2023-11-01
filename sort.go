package main

type sorter interface {
	sort([]int)
}

func swap(data []int, x, y int) {
	temp := data[x]
	data[x] = data[y]
	data[y] = temp
}

func isSorted(data []int) bool {
	for i, j := 0, 1; j < 20; i, j = i+1, j+1 {
		if data[i] > data[j] {
			return false
		}
	}

	return true
}

func bubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				swap(data, i, j)
			}
		}
	}
}

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data[j-1] > data[j]; j-- {
			swap(data, j, j-1)
		}
	}
}
