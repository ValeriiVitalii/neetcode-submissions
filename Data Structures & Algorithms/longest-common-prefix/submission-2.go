func longestCommonPrefix(strs []string) string {
    commonPrefix := ""

	for i:=0;len(strs[0]) > i;i++ {
		currentByte := strs[0][i]

		for _, str := range strs {
			if len(str) <= i || currentByte != str[i] {
				return commonPrefix
			}
		}

		commonPrefix = strs[0][:i+1]
	}

	return commonPrefix
}
