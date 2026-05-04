func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	words := make(map[byte]int, len(s))

	for i:=0;i<len(s);i++ {
		if _, ok := words[s[i]]; ok {
			words[s[i]]++
		} else {
			words[s[i]] = 1
		}

		if _, ok := words[t[i]]; ok {
			words[t[i]]--
		} else {
			words[t[i]] = -1
		}
	}

	for _, value := range words {
		if value != 0 {
			return false
		}
	}

	return true
}
