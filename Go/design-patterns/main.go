package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type QueueNode struct {
	node *TreeNode
	next *QueueNode
}

type Queue struct {
	first *QueueNode
	Size  int
}

func (q *Queue) Insert(node *TreeNode) {
	toInsert := &QueueNode{node: node}
	q.Size++
	if q.first == nil {
		q.first = toInsert
		return
	}

	toInsert.next = q.first
	q.first = toInsert
}

func (q *Queue) Pop() *TreeNode {
	if q.first == nil {
		return nil
	}

	if q.Size > 0 {
		q.Size--
	}

	toReturn := q.first.node
	q.first = q.first.next

	return toReturn
}

func (q *Queue) Empty() bool {
	return q.first != nil
}

func smallInQueue(root *TreeNode, q *Queue, k int) *TreeNode {
	if root.Left != nil {
		left := smallInQueue(root.Left, q, k)

		if left != nil {
			return left
		}
	}

	q.Insert(root)

	if q.Size == k {
		return q.Pop()
	}

	if root.Right != nil {
		right := smallInQueue(root.Right, q, k)
		if right != nil {
			return right
		}
	}

	return nil
}

func kthSmallest(root *TreeNode, k int) int {
	result := smallInQueue(root, &Queue{Size: 0}, k)
	if result != nil {
		return result.Val
	} else {
		return -1
	}
}

