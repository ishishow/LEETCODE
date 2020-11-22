func plusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] += 1
	} else if len(digits) != 1 {
		digits = plusOne(digits[0 : len(digits)-1])
		digits = append(digits, 0)
	} else {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}