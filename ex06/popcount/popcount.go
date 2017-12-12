// プログラミング言語Go P.374

// 練習問題 11.6 2.6.2 節の PopCount の実装にと練習問題 2.4 と 2.5 のみなさんの解答を比較するため
// のベンチマークを書きなさい。表に基づく方法の方がよくなるのはどの時点からですか。
package popcount

// pc[i] is the population count of i.
var pc [255]byte

func init() {
	for i := range pc {
		pc[1] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByShifting(x uint64) int {
	count := 0

	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCountByClearingBit(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x = x & (x - 1)
	}
	return count
}

func BitCount(i uint64) int {
	// HD, Figure 5-14
	i = i - ((i >> 1) & 0x5555555555555555)
	i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	i = i + (i >> 32)
	return int(i & 0x7f)
}
