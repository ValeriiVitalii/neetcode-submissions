func groupAnagrams(strs []string) [][]string {
	counts := make(map[[26]int][]string, len(strs))

	result := make([][]string, 0)

	for _, str := range strs {
		count := [26]int{}

		for _, r := range str {
			count[r - 'a']++
		}

		resultSlice, ok := counts[count]
		if ok {
			counts[count] = append(resultSlice, str)
		} else {
			counts[count] = []string{str}
		}
	}

	for _, resultSlice := range counts {
		result = append(result, resultSlice)
	}

	return result
}
