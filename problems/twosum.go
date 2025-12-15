package problems

// TwoSum returns indices of two numbers that add up to target.
// Returns (-1, -1) if no solution exists.
func TwoSum(nums []int, target int) (int, int) {
	seen := make(map[int]int)

	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return j, i
		}
		seen[v] = i
	}

	return -1, -1
}
