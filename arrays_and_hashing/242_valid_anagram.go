package arrays_and_hashing

// временная сложность O(n) - нужно пройтись один раз по обоим строкам
// память O(n) - нужно хранить мапу под одно слово
func isAnagram(s string, t string) bool {
	charsMap := make(map[rune]int)
	for _, char := range s {
		charsMap[char]++
	}
	for _, char := range t {
		if _, exists := charsMap[char]; !exists {
			return false
		}
		charsMap[char]--
	}

	for _, charsCount := range charsMap {
		if charsCount != 0 {
			return false
		}

	}

	return true
}
