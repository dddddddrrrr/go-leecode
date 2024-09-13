package combinationsum

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int
	var backtrack func(start, remain int)
	backtrack = func(start, remain int) {
		if remain == 0 {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remain {
				continue
			}
			current = append(current, candidates[i])
			backtrack(i, remain-candidates[i])
			current = current[:len(current)-1]
		}
	}
	backtrack(0, target)
	return result
}
