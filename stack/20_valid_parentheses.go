package stack

// временная сложность O(n) - обход строки
// сложность по памяти O(n) - хранение скобок в мапе
func isValid(s string) bool {
	// ({})
	// [{(]
	brackets := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	stack := make([]rune, 0)
	for _, br := range s {
		_, isOpen := brackets[br]
		if isOpen {
			stack = append(stack, br)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		lastOpenBracket := stack[len(stack)-1]
		checkCloseBracket := brackets[lastOpenBracket]
		if checkCloseBracket != br {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
