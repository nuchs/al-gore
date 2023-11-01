package main

import (
	"fmt"
	"time"
)

func runTests(tests []test, baseData []testData) {
	for _, t := range tests {
		fmt.Println(t.name)
		fmt.Println("---------------------------------------------")
		fmt.Printf("\t%11s  %11s  %11s\n", "Ascending", "Descending", "Random")
		for _, set := range baseData {
			asc := benchmark(t.sorter, set.asc)
			desc := benchmark(t.sorter, set.desc)
			rand := benchmark(t.sorter, set.rand)
			fmt.Printf("%s\t%8d µs  %8d µs  %8d µs\n", set.size, asc, desc, rand)
		}
		fmt.Println()
	}
}

func benchmark(test sorter, data []int) int {
	const iterations = 1
	total := 0
	for i := 0; i < iterations; i++ {
		result := runTest(test, data)
		if result == -1 {
			return -1
		}
		total += result
	}

	return total / iterations
}

func runTest(test sorter, baseData []int) int {
	data := copyData(baseData)
	start := time.Now()
	test(data)
	elapsed := time.Since(start)

	if !isSorted(data) {
		return -1
	} else {
		return int(elapsed.Microseconds())
	}
}
