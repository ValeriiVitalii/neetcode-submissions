func longestCommonPrefix(strs []string) string {
    commonPrefix := ""

	i := 0

main:
	for {
		var currentByte byte

		for _, str := range strs {
			if len(str) <= i {
				break main
			}

			if currentByte == 0 {
				currentByte = str[i]
			}

			if currentByte != str[i] {
				break main
			}
		}

		commonPrefix += string(currentByte)

		i++
	}

	return commonPrefix
}
