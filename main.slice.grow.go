package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Printf("initial: len: %d, cap: %2d, slice: %v, is nil: %t\n", len(s), cap(s), s, s == nil)

	for i := 0; i < 9; i++ {
		s = append(s, i)
		fmt.Printf("%d: len: %d, cap: %2d, slice: %v\n", i, len(s), cap(s), s)
	}
}
