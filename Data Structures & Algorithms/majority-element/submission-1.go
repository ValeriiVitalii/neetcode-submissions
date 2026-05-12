func majorityElement(nums []int) int {
    needApears := len(nums) / 2

	numsCount := make(map[int]int, 0)

	for _, num := range nums {
		numsCount[num]++

		if numsCount[num] > needApears {
			return num
		}
	}

	return 0
}
