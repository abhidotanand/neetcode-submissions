import "slices"
type MyStack struct {
	queue1 []int;
	queue2 []int;
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x)
	var rear = len(this.queue1) - 1
	this.queue2 = this.queue2[:0]
	for rear >= 0 {
		this.queue2 = append(this.queue2, this.queue1[rear])
		rear--
	}
}

func (this *MyStack) Pop() int {
	var tmp int
	if len(this.queue2) > 0 {
		tmp = this.queue2[0]
		this.queue2 = slices.Delete(this.queue2, 0, 1)
	}
	return tmp
}

func (this *MyStack) Top() int {
	return this.queue2[0]
}

func (this *MyStack) Empty() bool {
	if len(this.queue2) == 0 {
		return true
	}
	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
