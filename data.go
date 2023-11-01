package main

import "math/rand"

type generator int

const (
	ascending generator = iota
	descending
	random
)

func generateData(size int, gen generator) []int {
	data := make([]int, size)

	for i := 0; i < size; i++ {
		switch gen {
		case ascending:
			data[i] = i
		case descending:
			data[i] = size - i
		case random:
			data[i] = rand.Intn(10000)

		}
	}

	return data
}

func copyData(data []int) []int {
	dupe := make([]int, len(data))
	copy(data, dupe)

	return dupe
}
