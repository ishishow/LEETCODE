func searchInsert(nums []int, target int) int {
	count := 0

	for i, num := range nums {
		if num >= target {
			return i
		}
		count++
	}
	return count
}