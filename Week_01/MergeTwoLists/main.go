package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	now := &ListNode{}
	res := now
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			now.Next = l1
			l1 = l1.Next
		} else {
			now.Next = l2
			l2 = l2.Next
		}
		now = now.Next
	}
	if l1 != nil {
		now.Next = l1
	}
	if l2 != nil {
		now.Next = l2
	}
	return res.Next
}

func genRanSortedList(length int, maxVal int) *ListNode {
	if length == 0 {
		return nil
	}
	now := &ListNode{}
	res := now
	min := 0
	for i := 0; i < length; i++ {
		val := randInt(min, maxVal)
		// fmt.Println(val)
		min = val
		l := &ListNode{Val: val}
		now.Next = l
		now = l
	}
	return res.Next
}

func printList(l *ListNode) {
	for {
		if l == nil {
			break
		}
		fmt.Print(l.Val, ", ")
		l = l.Next
	}
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min >= max {
		return max
	}
	return rand.Intn(max-min) + min
}

func main() {
	l1 := genRanSortedList(10, 10000)
	l2 := genRanSortedList(15, 10000)

	printList(l1)
	fmt.Println("")
	printList(l2)
	fmt.Println("")
	printList(mergeTwoLists(l1, l2))
}
