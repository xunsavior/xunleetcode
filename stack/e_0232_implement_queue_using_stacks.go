package stack

/*
Company: Microsoft(6), Mathworks(3), Oracle(3), Apple(2). Bloomberg(2)

Implement the following operations of a queue using stacks.
push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.

Example:
MyQueue queue = new MyQueue();
queue.push(1);
queue.push(2);
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false

Notes:
You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue),
as long as you use only standard operations of a stack.
You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).
*/

// MyQueue ...
type MyQueue struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor232() MyQueue {
	return MyQueue{stack: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append([]int{x}, this.stack...)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	pop := this.stack[0]
	this.stack = this.stack[1:]
	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
