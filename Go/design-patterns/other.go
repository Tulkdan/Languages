package main

import "fmt"

type Tree struct {
	value       []int
	Left, Right *Tree
}

func (t *Tree) Insert(val int) {
	newValue := append(t.value, val)

	left := &Tree{value: newValue}
	right := &Tree{value: append([]int{}, t.value...)}

	fmt.Printf("Father %+v -> num %d -> toInsert %+v", t.value, val, newValue)

	t.Left = left
	t.Right = right
}

func findSomething(tree *Tree, val int) {
	if tree.Left == nil {
		tree.Insert(val)
		return
	}

	findSomething(tree.Left, val)
	findSomething(tree.Right, val)
}

func getLeafs(tree *Tree) [][]int {
	if tree.Left == nil {
		return [][]int{tree.value}
	}

	left := getLeafs(tree.Left)
	right := getLeafs(tree.Right)

	for _, r := range right {
		left = append(left, r)
	}

	return left
}

func printTree(tree *Tree) {
	if tree == nil {
		return
	}

	printTree(tree.Left)
	if tree.Left != nil {
		fmt.Printf("%+v", tree.Left.value)
	}

	printTree(tree.Right)
	if tree.Right != nil {
		fmt.Printf("%+v", tree.Right.value)
	}

	fmt.Printf("\n")
}

func subsets(nums []int) [][]int {
	tree := &Tree{value: []int{}}

	for _, num := range nums {
		findSomething(tree, num)
	}

	return getLeafs(tree)
}
