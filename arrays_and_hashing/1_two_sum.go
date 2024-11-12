package arrays_and_hashing

// временная сложность O(n) - нужно пройтись один раз массиву
// память O(n) - нужно хранить мапу значения
func twoSum(nums []int, target int) []int {
	//[2,7,11,15]
	//[2 => 0, ]
	seen := make(map[int]int)
	for k, v := range nums {
		secondArg := target - v
		if secondIdx, ok := seen[secondArg]; ok {
			return []int{k, secondIdx}
		}
		seen[v] = k
	}
	return []int{}
}
