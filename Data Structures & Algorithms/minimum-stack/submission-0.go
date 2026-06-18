type MinStack struct {
	s1 []int
	tmp []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.s1 = append(this.s1, val)
	if len(this.tmp) == 0 {
		this.tmp = append(this.tmp, val)
	} else {
		this.tmp = append(this.tmp, min(val, this.tmp[len(this.tmp) - 1]))
	}
}

func (this *MinStack) Pop() {
	this.s1 = this.s1[:len(this.s1)-1]
	this.tmp = this.tmp[:len(this.tmp)-1]
}

func (this *MinStack) Top() int {
	return this.s1[len(this.s1)-1]
}

func (this *MinStack) GetMin() int {
	return this.tmp[len(this.tmp)-1]
}
