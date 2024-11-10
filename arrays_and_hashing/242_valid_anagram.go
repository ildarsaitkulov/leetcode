package arrays_and_hashing

// временная сложность O(n) - нужно пройтись один раз строке и еще раз по мапе
// память O(n) - нужно хранить мапу под одно слово
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charsMap := make(map[rune]int)
	tlen := len(t)
	for i, char := range s {
		charsMap[char]++
		if i+1 > tlen {
			return false
		}
		charsMap[rune(t[i])]--
	}

	for _, charsCount := range charsMap {
		if charsCount != 0 {
			return false
		}

	}

	return true
}
