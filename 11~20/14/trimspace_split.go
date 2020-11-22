import (
	"strings"
)

func lengthOfLastWord(s string) int {
	strs := strings.Split(strings.TrimSpace(s), " ")
	return len(strs[len(strs)-1])
}