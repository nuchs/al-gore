package main

import (
	"testing"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Fuzz_Swap(f *testing.F) {
	data := generateData(30, random)
	f.Add(0, 0)
	f.Add(0, 1)
	f.Add(-1, -2)
	f.Fuzz(func(t *testing.T, first, second int) {
		ix := abs(first) % len(data)
		iy := abs(second) % len(data)
		x_orig := data[ix]
		y_orig := data[iy]

		swap(data, ix, iy)

		if data[ix] != y_orig || data[iy] != x_orig {
			t.Errorf(
				"expected data[%d]=%d, data[%d]=%d but got data[%d]=%d, data[%d]=%d",
				ix, x_orig,
				iy, y_orig,
				ix, data[ix],
				iy, data[iy])
		}
	})
}

func Test_DataGeneration(t *testing.T) {
	testCases := []struct {
		desc       string
		gen        generator
		comparison string
		check      func(int, int) bool
	}{
		{
			desc:       "Data in ascending order",
			gen:        ascending,
			comparison: ">",
			check:      func(x, y int) bool { return x <= y },
		},
		{
			desc:       "Data in descending order",
			gen:        descending,
			comparison: "<",
			check:      func(x, y int) bool { return x >= y },
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := generateData(20, tC.gen)
			for i, j := 0, 1; j < 20; i, j = i+1, j+1 {
				if !tC.check(got[i], got[j]) {
					t.Fatalf("Data not in correct order, got[%v]=%v %s got[%v]=%v", i, got[i], tC.comparison, j, got[j])
				}
			}
		})
	}
}

func Test_DataDuplication(t *testing.T) {
	testCases := []struct {
		desc  string
		check func([]int, []int) bool
	}{
		{
			desc: "Duplicate length should match",
			check: func(orig, dupe []int) bool {
				return len(orig) == len(dupe)
			},
		},
		{
			desc: "Duplicate content should match",
			check: func(orig, dupe []int) bool {
				for i := 0; i < len(orig); i++ {
					if orig[i] != dupe[i] {
						return false
					}
				}
				return true
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			orig := generateData(10, random)
			dupe := copyData(orig)
			if !tC.check(orig, dupe) {
				t.FailNow()
			}
		})
	}
}

func Test_SortChecking(t *testing.T) {
	testCases := []struct {
		name     string
		data     []int
		expected bool
	}{
		{
			name:     "Data is sorted",
			data:     generateData(20, ascending),
			expected: true,
		},
		{
			name:     "Data is not sorted",
			data:     generateData(20, descending),
			expected: false,
		},
	}
	for i, tc := range testCases {
		got := isSorted(tc.data)
		if got != tc.expected {
			t.Errorf(
				"\nTest %v: %v\ngot = %v, expected = %v\ndata = %v",
				i,
				tc.name,
				got,
				tc.expected,
				tc.data)
		}
	}
}
