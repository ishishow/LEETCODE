func removeDuplicates(nums []int) int {
	nw_index := 0
	for _, num := range nums {
		if nums[nw_index] != num {
			nw_index++
			nums[nw_index] = num
		}
	}

	return (nw_index + 1)
}