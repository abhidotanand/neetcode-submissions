type FreqStack struct {
    s map[int][]int
    m map[int]int
    maxCnt int
}


func Constructor() FreqStack {
    return FreqStack{
        make(map[int][]int),
        make(map[int]int),
        0,
    }
}


func (this *FreqStack) Push(val int)  {
    if _, ok := this.m[val]; ok {
        this.m[val]++
    } else {
        this.m[val] = 1
    }
    if this.m[val] > this.maxCnt {
        this.maxCnt = this.m[val]
    }
    this.s[this.m[val]] = append(this.s[this.m[val]], val)
}


func (this *FreqStack) Pop() int {
    tmp := this.s[this.maxCnt][len(this.s[this.maxCnt]) - 1]
    this.s[this.maxCnt] = this.s[this.maxCnt][:len(this.s[this.maxCnt])-1]
    if len(this.s[this.maxCnt]) == 0 {
        delete(this.s, this.maxCnt)
        this.maxCnt--
    }
    return tmp
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */