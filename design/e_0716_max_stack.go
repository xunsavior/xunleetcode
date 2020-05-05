package design

/*
Company: LinkedIn(13), Lyft(6), Amazon(3), Microsoft(2), Twitter(2)

Design a max stack that supports push, pop, top, peekMax and popMax.

push(x) -- Push element x onto stack.
pop() -- Remove the element on top of the stack and return it.
top() -- Get the element on the top.
peekMax() -- Retrieve the maximum element in the stack.
popMax() -- Retrieve the maximum element in the stack, and remove it. If you find more than one maximum elements, only remove the top-most one.

Example 1:
MaxStack stack = new MaxStack();
stack.push(5);
stack.push(1);
stack.push(5);
stack.top(); -> 5
stack.popMax(); -> 5
stack.top(); -> 1
stack.peekMax(); -> 5
stack.pop(); -> 1
stack.top(); -> 5

Note:
-1e7 <= x <= 1e7
Number of operations won't exceed 10000.
The last four operations won't be called when stack is empty.
*/

type MaxStack struct {
	stack    []int
	maxStack []int
}

/** initialize your data structure here. */
func Constructor716() MaxStack {
	return MaxStack{
		stack:    []int{},
		maxStack: []int{},
	}
}

func (this *MaxStack) Push716(x int) {
	this.stack = append(this.stack, x)
	if len(this.maxStack) == 0 {
		this.maxStack = append(this.maxStack, x)
	} else {
		peekMax := this.PeekMax716()
		if x > peekMax {
			this.maxStack = append(this.maxStack, x)
		} else {
			this.maxStack = append(this.maxStack, peekMax)
		}
	}
}

func (this *MaxStack) Pop716() int {
	pop, _ := this.stack[len(this.stack)-1], this.maxStack[len(this.maxStack)-1]
	this.stack, this.maxStack = this.stack[:len(this.stack)-1], this.maxStack[:len(this.maxStack)-1]
	return pop
}

func (this *MaxStack) Top716() int {
	return this.stack[len(this.stack)-1]
}

func (this *MaxStack) PeekMax716() int {
	return this.maxStack[len(this.maxStack)-1]
}

func (this *MaxStack) PopMax716() int {
	max := this.PeekMax716()
	buffer := []int{}
	val := this.Pop716()
	for val != max {
		buffer = append(buffer, val)
		val = this.Pop716()
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		this.Push716(buffer[i])
	}

	return max
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
