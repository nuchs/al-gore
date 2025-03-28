package main

import "math/rand"

type testData struct {
	size string
	asc  []int
	desc []int
	rand []int
}

type generator int

const (
	ascending generator = iota
	descending
	random
)

func generateDataSet() []testData {
	const small = 1000
	const medium = small * 100
	const large = medium * 100

	dataSet := make([]testData, 3)
	dataSet[0] = testData{
		"Small",
		generateData(small, ascending),
		generateData(small, descending),
		generateData(small, random),
	}
	dataSet[1] = testData{
		"Medium",
		generateData(medium, ascending),
		generateData(medium, descending),
		generateData(medium, random),
	}
	dataSet[2] = testData{
		"Large",
		generateData(large, ascending),
		generateData(large, descending),
		generateData(large, random),
	}

	return dataSet
}

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
