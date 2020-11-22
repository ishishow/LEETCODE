func climbStairs(n int) int {
	a := 1
	b := 1
	tmp := 0
	for i := 0; i < n; i++ {
		tmp = b
		b = a + b
		a = tmp
	}
	return a
}