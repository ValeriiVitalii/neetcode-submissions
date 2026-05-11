func groupAnagrams(strs []string) [][]string {
	counts := make(map[[26]int]int, len(strs))

	result := make([][]string, 0)

	for _, str := range strs {
		count := [26]int{}

		for _, r := range str {
			count[r - 'a']++
		}

		resultInx, ok := counts[count]
		if !ok {
			resultInx = len(result)

			counts[count] = resultInx
			
			result = append(result, []string{})
		}

		result[resultInx] = append(result[resultInx], str)
	}

	return result
}
