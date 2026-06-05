package main

import "testing"

// chapter3/sources/slice_benchmark_test.go
const sliceSize = 10000

func BenchmarkSliceInitWithoutCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := make([]int, 0)
		for i := 0; i < sliceSize; i++ {
			sl = append(sl, i)
		}
	}
}

func BenchmarkSliceInitWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := make([]int, 0, sliceSize)
		for i := 0; i < sliceSize; i++ {
			sl = append(sl, i)
		}
	}
}
