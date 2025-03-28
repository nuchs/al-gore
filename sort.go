package main

import (
	"context"
	"errors"
)

type sorter func([]int, context.Context) ([]int, error)

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

func bubbleSort(data []int, ctx context.Context) ([]int, error) {
	n := len(data)
	for i := range n {
		for j := 1; j < n-i; j++ {
			if data[j-1] > data[j] {
				swap(data, j-1, j)
			}
		}

		if err := ctx.Err(); err != nil {
			return nil, errors.New("Timed out")
		}
	}

	return data, nil
}

func insertionSort(data []int, ctx context.Context) ([]int, error) {
	n := len(data)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data[j-1] > data[j]; j-- {
			swap(data, j, j-1)
		}

		if err := ctx.Err(); err != nil {
			return nil, errors.New("Timed out")
		}
	}

	return data, nil
}
