import "strings"

func romanToInt(s string) int {
	var number int = 0
	var carry int = 0
	hashT := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for _, v := range strings.Split(s, "") {
		if carry < hashT[v] {
			number = number - (carry * 2) + hashT[v]
		} else {
			number = number + hashT[v]
		}
		carry = hashT[v]
	}
	return number
}