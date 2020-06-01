package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	if root.Left != nil {
		result = append(result, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, preorderTraversal(root.Right)...)
	}
	return result
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node1.Right = node2
	node2.Left = node3
	fmt.Println(preorderTraversal(node1))
}
