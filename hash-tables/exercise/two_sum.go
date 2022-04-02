package exercise

func TwoSum(nums []int, target int) [2]int {
	mapper := make(map[int]int)

	for i, num := range nums {
		if indices, ok := mapper[num]; ok {
			return [2]int{indices, i}
		}
		mapper[target-num] = i
	}

	return [2]int{0, 0}
}
