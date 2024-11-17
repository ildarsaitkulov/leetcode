package arrays_and_hashing

import "strings"

const delim = "#"

// временная сложность O(n)
// сложность по памяти O(n)
// где n длина строки захешированной
type Encoder struct {
}

func (e *Encoder) Encode(strs []string) string {
	b := strings.Builder{}
	for _, s := range strs {
		b.WriteString(s)
		b.WriteString(delim)
	}

	return b.String()
}

func (e *Encoder) Decode(s string) []string {
	runs := []rune(s)
	res := make([]string, 0)
	sb := strings.Builder{}
	for _, r := range runs {
		if string(r) != delim {
			sb.WriteRune(r)
			continue
		}
		res = append(res, sb.String())
		sb.Reset()
	}

	return res
}
