package arrays_and_hashing

// временная сложность O(n) - нужно пройтись один раз по обоим строкам
// память O(n) - нужно хранить мапу под одно слово
func isAnagram(s string, t string) bool {
	charsMap := make(map[rune]struct{})
	for _, char := range s {
		charsMap[char] = struct{}{}
	}
	for _, char := range t {
		if _, exists := charsMap[char]; !exists {
			return false
		}
	}

	return len(s) == len(t)
}
