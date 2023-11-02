package main

import (
	"context"
	"errors"
)

type sorter func([]int, context.Context) error

type test struct {
	name   string
	sorter sorter
}

func generateSortTests() []test {
	tests := make([]test, 2)

	tests[0] = test{name: "Bubble sort", sorter: bubbleSort}
	tests[1] = test{name: "Insertion sort", sorter: insertionSort}

	return tests
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

func bubbleSort(data []int, ctx context.Context) error {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				swap(data, i, j)
			}
		}

		if err := ctx.Err(); err != nil {
			return errors.New("Timed out")
		}
	}

	return nil
}

func insertionSort(data []int, ctx context.Context) error {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data[j-1] > data[j]; j-- {
			swap(data, j, j-1)
		}

		if err := ctx.Err(); err != nil {
			return errors.New("Timed out")
		}
	}

	return nil
}
