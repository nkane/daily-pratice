package leetcode

/*
	901: Online Stock Span
*/

type StockSpanner struct {
	monoStack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{}}
}

func (this *StockSpanner) Next(price int) int {
	res := 1
	for l := len(this.monoStack) - 1; l > -1 && this.monoStack[l][0] <= price; l-- {
		res += this.monoStack[l][1]
		this.monoStack = this.monoStack[:l]
	}
	this.monoStack = append(this.monoStack, [2]int{price, res})
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
