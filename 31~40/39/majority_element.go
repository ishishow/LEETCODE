func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}

	return map_max(m)
}

func map_max(m map[int]int) int {
	max := 0
	max_key := 0
	for key, value := range m {
		if max < value {
			max = value
			max_key = key
		}
	}
	return max_key
} 