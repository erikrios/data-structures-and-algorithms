package exercise

import "fmt"

func CountPairsWithDiff(nums []int, diff int) int {
	mapper := make(map[int]int)

	var count int
	for _, num := range nums {
		if v, ok := mapper[num]; ok {
			fmt.Println(num)
			count += v
		}

		mapper[num-diff] = mapper[num-diff] + 1
		mapper[num+diff] = mapper[num+diff] + 1
	}

	fmt.Println(mapper)

	return count
}
