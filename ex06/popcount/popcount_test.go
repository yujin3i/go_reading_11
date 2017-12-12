package popcount_test

import (
	"fmt"

	"github.com/yujin3i/go_reading_11/ex06/popcount"

	"testing"
)

func benchPopCount(b *testing.B, count int, v uint64) {
	var pc [255]byte

	for i := 0; i < b.N; i++ {
		for j := range pc {
			pc[j] = pc[j/2] + byte(j&1) // pcはグローバル変数
		}
		for j := 0; j < count; j++ {
			popcount.PopCount(v)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchmarks := []struct {
		Count int
		Value uint64
	}{
		{1, 0x1234567890ABCDEF},
		{10, 0x1234567890ABCDEF},
		{20, 0x1234567890ABCDEF},
		{30, 0x1234567890ABCDEF},
		{40, 0x1234567890ABCDEF},
		{50, 0x1234567890ABCDEF},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("%d, %X", bm.Count, bm.Value), func(b *testing.B) { benchPopCount(b, bm.Count, bm.Value) })
	}
}

func benchPopCountByShifting(b *testing.B, count int, v uint64) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			popcount.PopCountByShifting(v)
		}
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	benchmarks := []struct {
		Count int
		Value uint64
	}{
		{1, 0x1234567890ABCDEF},
		{10, 0x1234567890ABCDEF},
		{20, 0x1234567890ABCDEF},
		{30, 0x1234567890ABCDEF},
		{40, 0x1234567890ABCDEF},
		{50, 0x1234567890ABCDEF},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("%d, %X", bm.Count, bm.Value), func(b *testing.B) { benchPopCountByShifting(b, bm.Count, bm.Value) })
	}
}

func benchPopCountByClearingBit(b *testing.B, count int, v uint64) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			popcount.PopCountByClearingBit(v)
		}
	}
}

func BenchmarkPopCountByClearingBit(b *testing.B) {
	benchmarks := []struct {
		Count int
		Value uint64
	}{
		{1, 0x1234567890ABCDEF},
		{10, 0x1234567890ABCDEF},
		{20, 0x1234567890ABCDEF},
		{30, 0x1234567890ABCDEF},
		{40, 0x1234567890ABCDEF},
		{50, 0x1234567890ABCDEF},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("%d, %X", bm.Count, bm.Value), func(b *testing.B) { benchPopCountByClearingBit(b, bm.Count, bm.Value) })
	}
}

func benchBitCount(b *testing.B, count int, v uint64) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			popcount.BitCount(v)
		}
	}
}

func BenchmarkBitCount(b *testing.B) {
	benchmarks := []struct {
		Count int
		Value uint64
	}{
		{1, 0x1234567890ABCDEF},
		{10, 0x1234567890ABCDEF},
		{20, 0x1234567890ABCDEF},
		{30, 0x1234567890ABCDEF},
		{40, 0x1234567890ABCDEF},
		{50, 0x1234567890ABCDEF},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("%d, %X", bm.Count, bm.Value), func(b *testing.B) { benchBitCount(b, bm.Count, bm.Value) })
	}
}

/*
BenchmarkPopCount/1,_1234567890ABCDEF-4         	 5000000	       227 ns/op
BenchmarkPopCount/10,_1234567890ABCDEF-4        	 5000000	       250 ns/op
BenchmarkPopCount/20,_1234567890ABCDEF-4        	 5000000	       272 ns/op
BenchmarkPopCount/30,_1234567890ABCDEF-4        	 5000000	       299 ns/op
BenchmarkPopCount/40,_1234567890ABCDEF-4        	 5000000	       312 ns/op
BenchmarkPopCount/50,_1234567890ABCDEF-4        	 5000000	       335 ns/op
BenchmarkPopCountByShifting/1,_1234567890ABCDEF-4         	30000000	        53.9 ns/op
BenchmarkPopCountByShifting/10,_1234567890ABCDEF-4        	 3000000	       535 ns/op
BenchmarkPopCountByShifting/20,_1234567890ABCDEF-4        	 1000000	      1076 ns/op
BenchmarkPopCountByShifting/30,_1234567890ABCDEF-4        	 1000000	      1619 ns/op
BenchmarkPopCountByShifting/40,_1234567890ABCDEF-4        	 1000000	      2148 ns/op
BenchmarkPopCountByShifting/50,_1234567890ABCDEF-4        	  500000	      2707 ns/op
BenchmarkPopCountByClearingBit/1,_1234567890ABCDEF-4      	100000000	        20.2 ns/op
BenchmarkPopCountByClearingBit/10,_1234567890ABCDEF-4     	10000000	       198 ns/op
BenchmarkPopCountByClearingBit/20,_1234567890ABCDEF-4     	 5000000	       387 ns/op
BenchmarkPopCountByClearingBit/30,_1234567890ABCDEF-4     	 3000000	       581 ns/op
BenchmarkPopCountByClearingBit/40,_1234567890ABCDEF-4     	 2000000	       764 ns/op
BenchmarkPopCountByClearingBit/50,_1234567890ABCDEF-4     	 2000000	       959 ns/op
BenchmarkBitCount/1,_1234567890ABCDEF-4                   	2000000000	         0.81 ns/op
BenchmarkBitCount/10,_1234567890ABCDEF-4                  	500000000	         3.47 ns/op
BenchmarkBitCount/20,_1234567890ABCDEF-4                  	200000000	         6.28 ns/op
BenchmarkBitCount/30,_1234567890ABCDEF-4                  	200000000	         8.96 ns/op
BenchmarkBitCount/40,_1234567890ABCDEF-4                  	100000000	        11.7 ns/op
BenchmarkBitCount/50,_1234567890ABCDEF-4                  	100000000	        14.3 ns/op
*/
