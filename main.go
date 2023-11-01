package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Running sort algorithms")
	testData := [][]int{
		generateData(100, ascending),
		generateData(100, descending),
		generateData(100, random),
		generateData(1000, ascending),
		generateData(1000, descending),
		generateData(1000, random),
		generateData(10000, ascending),
		generateData(10000, descending),
		generateData(10000, random),
	}

	for _, baseData := range testData {
		data := copyData(baseData)
		start := time.Now()
		BubbleSort(data)
		taken := time.Since(start)
		if !isSorted(data) {
			fmt.Println("Failed to bubble sort data")
		} else {
			fmt.Printf("BubbleSort took %d nanoseconds\n", taken.Nanoseconds())
		}
	}
}
