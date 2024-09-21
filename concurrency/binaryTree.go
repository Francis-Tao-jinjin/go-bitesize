package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func create(v int) *Tree {
	return &Tree{Value: v}
}

func (t *Tree) insert(v int) {
	if v < t.Value {
		if t.Left == nil {
			t.Left = &Tree{Value: v}
		} else {
			t.Left.insert(v)
		}
	} else {
		if t.Right == nil {
			t.Right = &Tree{Value: v}
		} else {
			t.Right.insert(v)
		}
	}
}

func (t *Tree) preOrderTraversal() []int {
	if t == nil {
		return []int{}
	}
	var result []int
	var traversal func(node *Tree)
	traversal = func(node *Tree) {
		if node == nil {
			return
		}
		result = append(result, node.Value)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(t)
	return result
}

func (t *Tree) inOrderTraversal() []int {
	if t == nil {
		return []int{}
	}
	var result []int
	var traversal func(node *Tree)
	traversal = func(node *Tree) {
		if node == nil {
			return
		}
		traversal(node.Left)
		result = append(result, node.Value)
		traversal(node.Right)
	}
	traversal(t)
	return result
}

func (t *Tree) bfs() []int {
	var result []int
	var queue []*Tree
	queue = append(queue, t)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func (t *Tree) levelOrderTraversal() [][]int {
	var result [][]int
	var queue []*Tree
	queue = append(queue, t)
	for len(queue) > 0 {
		var level []int
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Value)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

func buildTree(inOrderList []int, preOrderList []int) *Tree {
	if len(inOrderList) == 0 {
		return nil
	}
	root := &Tree{Value: preOrderList[0]}
	inOrderRootIndex := slices.Index(inOrderList, preOrderList[0])
	leftInOrderList := inOrderList[:inOrderRootIndex]
	rightInOrderList := inOrderList[inOrderRootIndex+1:]
	leftPreOrderList := preOrderList[1 : 1+inOrderRootIndex]
	rightPreOrderList := preOrderList[1+inOrderRootIndex:]

	root.Left = buildTree(leftInOrderList, leftPreOrderList)
	root.Right = buildTree(rightInOrderList, rightPreOrderList)
	return root
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	var walk func(node *Tree)
	walk = func(node *Tree) {
		if node == nil {
			return
		}
		ch <- node.Value
		walk(node.Left)
		walk(node.Right)
	}
	walk(t)
	close(ch)
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		v2, ok := <-ch2
		fmt.Printf("v1: %v, v2: %v\n", v1, v2)
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func TryBinaryTree() {

	treeA := buildTree([]int{1, 1, 2, 3, 5, 8, 13}, []int{3, 1, 1, 2, 8, 5, 13})
	treeB := buildTree([]int{1, 1, 2, 3, 5, 8, 13}, []int{8, 3, 1, 1, 2, 5, 13})
	treeC := buildTree([]int{1, 1, 2, 3, 5, 8, 13}, []int{8, 3, 1, 1, 2, 5, 13})

	fmt.Println(treeA.levelOrderTraversal())
	fmt.Println(treeB.levelOrderTraversal())

	fmt.Println(Same(treeB, treeC))
}
