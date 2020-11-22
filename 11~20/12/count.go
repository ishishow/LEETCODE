import "strconv"

func countAndSay(n int) string {
	seq := "1"
	for i := 1; i < n; i++ {
		seq = next(seq)
	}
	return seq
}

func next(seq string) string {
	i := 0
	nextSeq := ""
	for i < len(seq) {
		count := 1
		for i < len(seq)-1 && seq[i] == seq[i+1] {
			count++
			i++
		}
		nextSeq += strconv.Itoa(count) + string(seq[i])
		i++
	}
	return nextSeq
}