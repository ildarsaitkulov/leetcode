package arrays_and_hashing

// временная сложность O(n) - нужно пройтись один раз по массиву
// память O(n) - в худшем случае придется сохранить все в мапу
func containsDuplicate(nums []int) bool {
	existsMap := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := existsMap[num]; exists {
			return true
		}
		existsMap[num] = struct{}{}
	}

	return false
}
