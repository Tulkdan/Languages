package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
    permutations := [][]int{}

    var backtrack func([]int, []bool)
    backtrack = func (curr []int, pick []bool) {
        if len(curr) == len(nums) {
            permutations = append(permutations, append([]int{}, curr...))
            return
        }

        for j, val := range nums {
            if pick[j] {
                continue
            }

            pick[j] = true
            backtrack(append(curr, val), pick)
            pick[j] = false
        }
    }

    backtrack([]int{}, make([]bool, len(nums)))

    return permutations
}


func subsetWithDup(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}

    var backtrack func([]int, int)
    backtrack = func(curr []int, i int) {
        result = append(result, append([]int{}, curr...))

        for j := i; j < len(nums); j++ {
            if j > i && nums[j] == nums[j - 1] {
                continue
            }
            backtrack(append(curr, nums[j]), j+1)
        }
    }

    backtrack([]int{}, 0)

    return result
}


func main() {
	/**
	a := exist([][]byte{
		{'A', 'B', 'C', 'D'},
		{'S', 'A', 'A', 'T'},
		{'A', 'C', 'A', 'E'},
	}, "CAT")
	fmt.Println(a)

	a = exist([][]byte{
		{'A', 'B', 'C', 'D'},
		{'S', 'A', 'A', 'T'},
		{'A', 'C', 'A', 'E'},
	}, "BAT")
	fmt.Println(a)

	*/
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCB"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist([][]byte{{'A'}}, "A"))
}

func exist(board [][]byte, word string) bool {
	hasSeen := make([][]bool, len(board))
	for i, _ := range hasSeen {
		hasSeen[i] = make([]bool, len(board[0]))
	}

	var backtracking func(int, int, int) bool
	backtracking = func(x, y, l int) bool {
		if l == len(word) {
			return true
		}
		fmt.Printf("letter %q at X->%d Y->%d\n", word[l], x, y)

		if x < 0 || y < 0 ||
			x >= len(board[0]) || y >= len(board) ||
			board[y][x] != word[l] || hasSeen[y][x] {
			return false
		}

		fmt.Printf("Found letter %q at X -> %d Y -> %d\n", word[l], x, y)
		hasSeen[y][x] = true
		res := backtracking(x+1, y, l+1) ||
			backtracking(x, y+1, l+1) ||
			backtracking(x-1, y, l+1) ||
			backtracking(x, y-1, l+1)
		hasSeen[y][x] = false

		return res
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if backtracking(x, y, 0) {
				return true
			}
		}
	}
	return false
}
