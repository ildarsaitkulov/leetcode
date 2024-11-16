package arrays_and_hashing

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	n := len(nums)
	freq := make([][]int, n+1)
	for num, c := range count {
		freq[c] = append(freq[c], num)
	}

	res := []int{}
	for i := n; i > 0 && len(res) < k; i-- {
		if len(freq[i]) > 0 {
			res = append(res, freq[i]...)
		}
	}

	return res[:k]
}
