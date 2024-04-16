package minpossibleint1505

import (
	"container/list"
)

// The FenwickTree struct represents a data structure used for efficient range queries and updates on
// an array.
// @property {int} size - The `size` property in the `FenwickTree` struct represents the length of the
// array for which the Fenwick tree is being constructed. It is denoted by `n` in the context of
// Fenwick trees.
// @property {[]int} bitTree - The `bitTree` property in the `FenwickTree` struct is used to store the
// sum of ranges. It is a binary indexed tree (also known as a Fenwick tree) data structure that allows
// for efficient calculation of prefix sums and updates in an array of numbers. The `bit
type FenwickTree struct {
	size    int   // n: length of the array
	bitTree []int // bit: store the sum of range
}

// The NewFenwickTree function in Go initializes and returns a new FenwickTree data structure with a
// specified size.
func NewFenwickTree(n int) *FenwickTree {
	ft := FenwickTree{
		size:    n + 1,
		bitTree: make([]int, n+1),
	}
	return &ft
}

// This `CountElementsLeft` method in the `FenwickTree` struct is used to calculate the number of elements already shifted in
// the num located left from the index `i`.
func (ft *FenwickTree) CountElementsLeft(i int) int {
	sum := 0
	for i > 0 {
		sum += ft.bitTree[i]
		i -= (i & -i) // flip the last set bit
	}
	return sum
}

// The `Add` method in the `FenwickTree` struct is used to update the Fenwick tree by adding a value to
// a specific index `i` and propagating the changes through the tree structure. Here is a breakdown of
// what the method does:
func (ft *FenwickTree) Add(i, inc int) {
	for i < ft.size {
		ft.bitTree[i] += inc
		i += (i & -i) // add last set bit
	}
}

// The `minIntegerBit` function in Go takes a string of digits and an integer `k`, and rearranges the
// digits to form the smallest possible number by performing a specific number of left shifts on the
// digits.
func minIntegerBit(num string, k int) string {
	n := len(num)
	// Map of digits [0-9] and list of positions to the each digit
	positions := make([]*list.List, 10)
	// Init Linked list for each digit
	for i := 0; i <= 9; i++ {
		positions[i] = list.New()
	}
	// Go through num and save each position corresponding to digit
	for i := 0; i < n; i++ {
		d := num[i] - '0'
		positions[d].PushBack(i + 1) // indexing starts from 1 in Binary Indexed Tree
	}

	ft := NewFenwickTree(n)
	answer := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		for d := 0; d <= 9; d++ {
			posEl := positions[d].Front()
			if posEl != nil {
				position := posEl.Value.(int)
				leftShifts := ft.CountElementsLeft(position)
				if position-1-leftShifts <= k {
					answer = append(answer, byte(d+'0'))
					ft.Add(position, 1)
					k -= (position - 1 - leftShifts)
					positions[d].Remove(posEl)
					break // element added, jump to the next iteration
				}
			}
		}
	}
	return string(answer)
}
