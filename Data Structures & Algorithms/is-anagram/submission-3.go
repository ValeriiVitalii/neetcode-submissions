func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	words := [26]int{}

	for i:=0;i<len(s);i++ {
		words[s[i] - 'a']++
		words[t[i] - 'a']--
	}

	for _, value := range words {
		if value != 0 {
			return false
		}
	}

	return true
}
