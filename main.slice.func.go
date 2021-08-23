package main

import (
	"fmt"
)

func grow(s []int) {
	s = append(s, 2)
	fmt.Printf("%6s: len: %d, cap: %d, slice: %v\n", "inside", len(s), cap(s), s)
}

func main() {
	s := make([]int, 0, 3)
	s = append(s, []int{0, 1}...)

	fmt.Printf("%6s: len: %d, cap: %d, slice: %v\n", "before", len(s), cap(s), s)

	grow(s)

	fmt.Printf("%6s: len: %d, cap: %d, slice: %v\n", "after", len(s), cap(s), s)
}
