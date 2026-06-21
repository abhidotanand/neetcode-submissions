type StockSpanner struct {
	stack [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	if len(this.stack) == 0 {
        this.stack = append(this.stack, []int{price, 1})
        return 1
    } else {
        if price < this.stack[len(this.stack)-1][0] {
            this.stack = append(this.stack, []int{price, 1})
            return 1
        } else {
            tmp := 1
            for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1][0] {
                tmp += this.stack[len(this.stack)-1][1]
                this.stack = this.stack[:len(this.stack)-1]
            }
            this.stack = append(this.stack, []int{price, tmp})
            return tmp
        }
    }
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 