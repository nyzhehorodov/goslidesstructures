package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Printf("len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)

	s = s[2:4]
	fmt.Printf("len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)

	s = s[:cap(s)]
	fmt.Printf("len: %d, cap: %d, slice: %v\n", len(s), cap(s), s)
}
