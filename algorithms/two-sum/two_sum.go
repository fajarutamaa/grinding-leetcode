package twosum

func TwoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, ok := idxMap[complement]; ok {
			return []int{idx, i}
		}
		idxMap[num] = i
	}
	return nil
}
