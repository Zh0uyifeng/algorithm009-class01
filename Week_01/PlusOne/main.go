package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	a1 := []int{1, 2, 3}
	a2 := []int{}
	a3 := []int{9, 9, 9}
	a4 := []int{1, 0, 9}
	fmt.Println(plusOne(a1))
	fmt.Println(plusOne(a2))
	fmt.Println(plusOne(a3))
	fmt.Println(plusOne(a4))
}
