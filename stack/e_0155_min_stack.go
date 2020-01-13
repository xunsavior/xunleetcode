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

type node155 struct {
	Val  int
	Next *node155
	Min  int
}

// MinStack ...
type MinStack struct {
	Head *node155
}

// Constructor155 ...
func Constructor155() MinStack {
	return MinStack{}
}

func (this *MinStack) Push155(x int) {
	newStack := &node155{Val: x}
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

func (this *MinStack) Pop155() {
	this.Head = this.Head.Next
}

func (this *MinStack) Top155() int {
	return this.Head.Val
}

func (this *MinStack) GetMin() int {
	return this.Head.Min
}
