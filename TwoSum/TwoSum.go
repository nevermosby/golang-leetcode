package twosum

func twoSum(nums []int, target int) []int {
	mapv := make(map[int]int)
	for i, v := range nums {
		if val, ok := mapv[target-v]; ok {
			return []int{i, val}
		}
		mapv[v] = i
	}
	return nil
}
