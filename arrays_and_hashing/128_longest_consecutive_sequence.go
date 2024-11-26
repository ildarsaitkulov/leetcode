package arrays_and_hashing

// временная сложность O(n) - из-за того что есть проверка начала последовательности
// сложность по памяти O(n)
func longestConsecutive(nums []int) int {
	numsMap := make(map[int]struct{})
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	maxLen := 0
	for num := range numsMap {
		prev := num - 1
		if _, exists := numsMap[prev]; exists {
			continue
		}

		currentLen := 1
		for i := num + 1; ; i++ {
			if _, existsNext := numsMap[i]; !existsNext {
				break
			}
			currentLen++
		}
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
