package two_pointers

import (
	"regexp"
	"strings"
)

// Временная и по памяти сложность O(n)
func isPalindrome(s string) bool {
	s = prepareStr(s)
	runes := []rune(s)
	left := 0
	right := len(runes) - 1

	for left <= right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func prepareStr(s string) string {
	r := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleaned := r.ReplaceAllString(s, "")
	return strings.ToLower(cleaned)
}
