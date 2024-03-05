package stackmakequeue

type MyQueue struct {
	stackIn 	[]int
	stackOut	[]int
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) inToOut(){
	for len(this.stackIn) > 0 {
		this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
		this.stackIn = this.stackIn[:len(this.stackIn)-1]
	}
}


func (this *MyQueue) Pop() int {
	if len(this.stackOut) == 0 {
		this.inToOut()
	}
	result := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	return result
}


func (this *MyQueue) Peek() int {
	if len(this.stackOut) == 0 {
		this.inToOut()
	}
	result := this.stackOut[len(this.stackOut)-1]
	return result
}


func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
