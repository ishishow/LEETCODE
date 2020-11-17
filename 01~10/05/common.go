func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	max := len(strs[0])

	for i := 0; i < max; i++ {
		b := strs[0][i]
		for _, str := range strs[1:] {
			if i == len(str) || b != str[i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}