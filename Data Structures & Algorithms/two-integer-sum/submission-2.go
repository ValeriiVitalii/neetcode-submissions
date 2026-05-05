func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int, len(nums))

	for i, num := range nums {
		difference := target - num

		if value, ok := numsMap[difference]; ok {
			return []int{value, i}
		} else {
			numsMap[num] = i
		}
	}

	return []int{}
}
