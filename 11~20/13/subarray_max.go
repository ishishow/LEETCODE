func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
	}
	return max_num(nums)
}

func max_num(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}