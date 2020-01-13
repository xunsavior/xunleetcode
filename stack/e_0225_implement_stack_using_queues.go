package stack

/*
Company: Microsoft(5), Amazon(2), Citadel(2); Twilio(3), Mathworks(2)

Implement the following operations of a stack using queues.
push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
empty() -- Return whether the stack is empty.

Example:
MyStack stack = new MyStack();
stack.push(1);
stack.push(2);
stack.top();   // returns 2
stack.pop();   // returns 2
stack.empty(); // returns false

Notes:
You must use only standard operations of a queue -- which means only push to back,
peek/pop from front, size, and is empty operations are valid.

Depending on your language, queue may not be supported natively. You may simulate
a queue by using a list or deque (double-ended queue), as long as you use only
standard operations of a queue.

You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).
*/

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor225() MyStack {
	return MyStack{queue: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	n := len(this.queue)
	pop := (this.queue)[n-1]
	(this.queue) = (this.queue)[:n-1]
	return pop
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len((this.queue)) == 0
}
