package exercise

func MostFrequent(nums []int) int {
	mapper := make(map[int]int)

	for _, num := range nums {
		count := mapper[num]
		mapper[num] = count + 1
	}

	var frequent int
	for k, v := range mapper {
		if mapper[frequent] < v {
			frequent = k
		}
	}

	return frequent
}
