package main

import (
	"errors"
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
			fmt.Printf("%s\t%8v  %8v  %8v\n", set.size, asc, desc, rand)
		}
		fmt.Println()
	}
}

func benchmark(test sorter, data []int) string {
	const iterations = 1
	total := 0
	for i := 0; i < iterations; i++ {
		if result, err := runTest(test, data); err != nil {
			return err.Error()
		} else {
			total += result
		}
	}

	return fmt.Sprintf("%d µs", total/iterations)
}

func runTest(test sorter, baseData []int) (int, error) {
	data := copyData(baseData)
	start := time.Now()
	test(data)
	elapsed := time.Since(start)

	if !isSorted(data) {
		return -1, errors.New("Fail")
	} else {
		return int(elapsed.Microseconds()), nil
	}
}
