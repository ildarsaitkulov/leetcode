package arrays_and_hashing

import "sort"

// n кол-во строк, m - среднее кол-во символов
// сложность по времени O(n * m log m), O(n) нужно для того чтобы пройтись по исходному слайсу и O(m*log*m) для сортировки строк
// сложность по памяти O(n * m) - столько нужно для хранения мапы из n строк и длиной k
func groupAnagrams(strs []string) [][]string {
	grouped := make(map[string][]string)
	for _, s := range strs {
		a := []rune(s)
		sort.Slice(a, func(i, j int) bool {
			return a[i] > a[j]
		})

		grouped[string(a)] = append(grouped[string(a)], s)
	}

	res := make([][]string, 0, len(grouped))

	for _, items := range grouped {
		res = append(res, items)
	}

	return res
}
