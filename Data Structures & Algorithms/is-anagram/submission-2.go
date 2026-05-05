func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	words := make(map[byte]int, len(s))

	for i:=0;i<len(s);i++ {
		words[s[i]]++

		words[t[i]]--
	}

	for _, value := range words {
		if value != 0 {
			return false
		}
	}

	return true
}
