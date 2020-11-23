type MinStack struct {
	vals [][]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{vals: [][]int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.vals) == 0 {
		this.vals = [][]int{{x, x}}
	} else {
		this.vals = append(this.vals, []int{x, min(x, this.GetMin())})
	}
}

func (this *MinStack) Pop() {
	this.vals = this.vals[:len(this.vals)-1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.vals[len(this.vals)-1][1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}