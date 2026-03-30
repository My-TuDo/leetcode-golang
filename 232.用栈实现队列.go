/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

// 辅助函数：如果outStack空了，就从inStack中进
func (this *MyQueue) move() {
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			// 拿到inStack的栈顶
			val := this.inStack[len(this.inStack)-1]
			// 弹出inStack的栈顶
			this.inStack = this.inStack[:len(this.inStack)-1]
			// 压入outStack
			this.outStack = append(this.outStack, val)
		}
	}
}

func (this *MyQueue) Pop() int {
	this.move()
	val := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return val
}

func (this *MyQueue) Peek() int {
	this.move()
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

