func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int, len(nums))

	for i, num := range nums {
		difference := target - num

		if value, ok := indexMap[difference]; ok {
			return []int{value, i}
		} 

		indexMap[num] = i
		
	}

	return nil
}
