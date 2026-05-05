func hasDuplicate(nums []int) bool {
	numbers := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := numbers[num]; ok {
			return true
		}

		numbers[num] = struct{}{}
	}

	return false
}
