func removeElement(nums []int, val int) int {

	nw_index := 0
	for _, num := range nums {
		if num != val {
			nums[nw_index] = num
			nw_index++
		}
	}

	return nw_index
}