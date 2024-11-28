package two_pointers

// Временная сложность O(n)
// По памяти O(1) дополнительной памяти не нужно, только под ответ
func twoSum(numbers []int, target int) []int {
	//[-5,-3,0,2,4,6,8]
	//     ^         ^
	//5

	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[right] + numbers[left]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}
