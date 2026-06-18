type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	var rear2 int = len(this.s2) - 1
	for rear2 >= 0 {
		this.s1 = append(this.s1, this.s2[rear2])
		rear2--
	}

	this.s1 = append(this.s1, x)

	var rear1 int = len(this.s1) - 1
	for rear1 >= 0 {
		this.s2 = append(this.s2, this.s1[rear1])
		rear1--
	}
	this.s1 = nil
}

func (this *MyQueue) Pop() int {
	var tmp int = this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return tmp
}

func (this *MyQueue) Peek() int {
	return this.s2[len(this.s2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
