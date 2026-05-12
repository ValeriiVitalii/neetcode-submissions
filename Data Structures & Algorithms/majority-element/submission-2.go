func majorityElement(nums []int) int {
    candidat := nums[0]

	count := 0

	for _, num := range nums {
		if count == 0 {
			candidat = num
		}

		if num == candidat {
			count ++
		} else {
			count--
		}
	}

	return candidat
}
