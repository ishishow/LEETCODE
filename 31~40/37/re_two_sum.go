func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		idx, ok := m[target-numbers[i]]
		if ok {
			return []int{idx + 1, i + 1}
		}
		m[numbers[i]] = i
	}
	return nil
}