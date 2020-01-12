package stack

/*
Company: Amazon(17), Bloomberg(9), Microsoft(8), Apple(4)

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/

type StackNode struct {
	Val  int
	Next *StackNode
	Min  int
}

type MinStack struct {
	Head *StackNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	newStack := &StackNode{Val: x}
	if this.Head == nil {
		newStack.Min = x
		this.Head = newStack
	} else {
		if x < this.Head.Min {
			newStack.Min = x
		} else {
			newStack.Min = this.Head.Min
		}
		newStack.Next, this.Head = this.Head, newStack
	}
}

func (this *MinStack) Pop() {
	this.Head = this.Head.Next
}

func (this *MinStack) Top() int {
	return this.Head.Val
}

func (this *MinStack) GetMin() int {
	return this.Head.Min
}
