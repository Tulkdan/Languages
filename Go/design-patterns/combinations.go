package main

import "sort"

func dfs(candidates []int, target int, i int, total int, curr []int, res [][]int) [][]int {
	if total == target {
		return append(res, curr)
	}

	if total > target || i >= len(candidates) {
		return res
	}

	var newCurr = make([]int, len(curr)+1)
	copy(newCurr, append(curr, candidates[i]))
	newRes := dfs(candidates, target, i, total+candidates[i], newCurr, res)

	nextRes := dfs(candidates, target, i+1, total, curr, res)

	return append(newRes, nextRes...)
}

func combinationSum(candidates []int, target int) [][]int {
	return dfs(candidates, target, 0, 0, []int{}, [][]int{})
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := [][]int{}

	var backtrace func(int, int, []int)
	backtrace = func(i int, val int, curr []int) {
		if val == target {
			result = append(result, append([]int{}, curr...))
			return
		}

		if val > target {
			return
		}

		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j] == candidates[j - 1] {
				continue
			}

			backtrace(j+1, val + candidates[j], append(curr, candidates[j]))
		}
	}

	backtrace(0, 0, []int{})

	return result
}

