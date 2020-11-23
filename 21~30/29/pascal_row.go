func getRow(rowIndex int) []int {
	ReArray := make([]int, rowIndex+1)
	ReArray[rowIndex], ReArray[0] = 0, 1
	for i := 0; i < rowIndex+1; i++ {
		for j := i; j >= 1; j-- {
			ReArray[j] += ReArray[j-1]
		}
	}
	return ReArray
}