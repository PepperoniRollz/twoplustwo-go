package generator

import (
	c "github.com/pepperonirollz/twoplustwo-go/constants"
)

var flushCount int = 0

func eval5HandFast(c1, c2, c3, c4, c5 int) int {
	q := (c1 | c2 | c3 | c4 | c5) >> 16
	var s int16

	if c1&c2&c3&c4&c5&0xf000 != 0 {

		return int(c.FLUSHES[q])

	}
	s = c.UNIQUE_5[q]
	if s != 0 {
		return int(s)
	}

	return int(c.HASH_VALUES[findFast(uint32((c1&0xff)*(c2&0xff)*(c3&0xff)*(c4&0xff)*(c5&0xff)))])
}

func findFast(u uint32) uint32 {
	var a, b, r uint32
	u += 0xe91aaa35
	u ^= u >> 16
	u += u << 8
	u ^= u >> 4
	b = (u >> 8) & 0x1ff
	a = (u + (u << 2)) >> 19
	r = a ^ uint32(c.HASH_ADJUST[b])
	return r
}

func eval7hand(hand [8]int) int {
	best := 9999
	subhand := make([]int, 5)

	for i := 0; i < 21; i++ {
		for j := 0; j < 5; j++ {
			subhand[j] = hand[c.PERM_7[i][j]]
		}
		q := eval5hand(subhand)
		if q < best {
			best = q
		}
	}

	return best
}

func eval5hand(hand []int) int {
	c1 := hand[0]
	c2 := hand[1]
	c3 := hand[2]
	c4 := hand[3]
	c5 := hand[4]

	return eval5HandFast(c1, c2, c3, c4, c5)
}
