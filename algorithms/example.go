package main

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	// used[i] 表示 candidates[i] 是否在本次递归中被使用
	var backtracking func(path []int, used []bool, idx int, rest int)
	backtracking = func(path []int, used []bool, idx, rest int) {
		if rest < 0 {
			return
		}
		if rest == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := idx; i < len(candidates); i++ {
			// 如果candidates[i] == candidates[i - 1] 并且 used[i - 1] == false，就说明：
			// 前一个树枝，使用了candidates[i - 1]，也就是说同一树层使用过candidates[i - 1]。
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, candidates[i])
			backtracking(path, used, idx+1, rest-candidates[i])
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking(make([]int, 0), make([]bool, len(candidates)), 0, target)
	return res
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var backtracking func(path []int, used []bool)
	backtracking = func(path []int, used []bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtracking(path, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking(make([]int, 0), make([]bool, len(nums)))
	return res
}
