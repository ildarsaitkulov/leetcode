package arrays_and_hashing

// O(n)
func productExceptSelf(nums []int) []int {
	//[1,  2,  3,  4]
	//[1,  2,  6,  24] left
	//[24, 24, 12, 4] right
	//[24, 12, 8,  6] result

	rightProduct := make([]int, len(nums), len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]
		if i == len(nums)-1 {
			rightProduct[i] = n
		} else {
			rightProduct[i] = rightProduct[i+1] * n
		}
	}

	result := make([]int, len(nums), len(nums))

	leftProduct := make([]int, len(nums), len(nums))
	for i, n := range nums {
		if i == 0 {
			leftProduct[i] = n
			result[i] = rightProduct[i+1]
		} else {
			leftProduct[i] = leftProduct[i-1] * n
			if i == len(nums)-1 {
				result[i] = leftProduct[i-1]
			} else {
				result[i] = leftProduct[i-1] * rightProduct[i+1]
			}
		}
	}
	return result
}
