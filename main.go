package main

import (
	"fmt"
	"github.com/m1gwings/treedrawer/tree"
	"github.com/oze4/jslice"
)

func main() {
	n := 7

	results := allPossibleFBT(n)

	for _, result := range results {
		printTreeNode(result)
	}
}

// TreeNode is essentially a doubly-linked list.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// allPossibleFBT prints all possible full binary trees up to `n` nodes.
func allPossibleFBT(n int) []*TreeNode {
	head := []*TreeNode{}
	// If n is even it's not possible.
	if n%2 == 0 {
		return head
	}
	if n == 1 {
		return append(head, &TreeNode{})
	}

	for i := 1; i < n; i += 2 {
		lefts := allPossibleFBT(i)
		rights := allPossibleFBT(n - 1 - i)

		for _, left := range lefts {
			for _, right := range rights {
				head = append(head, &TreeNode{Val: 0, Left: left, Right: right})
			}
		}
	}
	return head
}

// printTreeNode prints out a *TreeNode (binary tree) to console.
// Uses bfs to generate the `*tree.Tree`
func printTreeNode(head *TreeNode) *tree.Tree {
	if head == nil {
		return nil
	}

	treeRoot := tree.NewTree(tree.NodeInt64(head.Val))
	// Assign head "children" as the head itself.
	// This is due to the first pass where our TreeNode head
	// is the only thing in the queue.
	treeLeftChild := treeRoot
	treeRightChild := treeRoot

	QUEUE := []*TreeNode{head}

	for len(QUEUE) > 0 {
		size := len(QUEUE)

		for i := 0; i < size; i++ {
			node := jslice.Shift(&QUEUE)

			// Alternate between left and right children.
			// We expect left child first (and to alternate from there) in the queue.
			treeChild := treeLeftChild
			if i%2 == 1 {
				treeChild = treeRightChild
			}

			if node.Left != nil {
				jslice.Push(&QUEUE, node.Left)
				treeLeftChild = treeChild.AddChild(tree.NodeInt64(node.Left.Val))
			}
			if node.Right != nil {
				jslice.Push(&QUEUE, node.Right)
				treeRightChild = treeChild.AddChild(tree.NodeInt64(node.Right.Val))
			}
		}
	}

	fmt.Println(treeRoot)
	return treeRoot
}
