package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	res := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			res = append(res, value)
			res = append(res, i)
		}
		m[k] = i
	}
	return res

}

func main() {

	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

}
