package samestructure

func SameStructure(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	counter := make(map[rune]int)

	for _, ch := range s1 {
		counter[ch]++
	}

	for _, ch := range s2 {
		counter[ch]--
		if counter[ch] < 0 {
			return false
		}
	}
	return true
}