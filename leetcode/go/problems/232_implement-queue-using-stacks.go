package problems

//MyQueue implement queue using stack
type MyQueue struct {
	input  []int
	output []int
}

//Constructor Initialize your data structure here.
func ConstructorMyQueue() MyQueue {
	input := make([]int, 0)
	output := make([]int, 0)
	return MyQueue{input, output}

}

//Push Push element x to the back of queue.
func (t *MyQueue) Push(x int) {
	t.input = append(t.input, x)

}

//Pop Removes the element from in front of queue and returns that element.
func (t *MyQueue) Pop() int {
	defer func() {
		if len(t.output) > 0 {
			t.output = t.output[:len(t.output)-1]
		}
	}()
	return t.Peek()
}

//Peek  Get the front element.
func (t *MyQueue) Peek() int {
	// 只需要保证，输入的元素总是跟在前面的输入元素的后面，而输出元素总是最早输入的那个元素即可
	// 只有在「输出栈」为空的时候，才发生一次性的「倒腾」
	if len(t.output) == 0 {
		for len(t.input) > 0 {
			t.output = append(t.output, t.input[len(t.input)-1])
			t.input = t.input[:len(t.input)-1]
		}
	}
	return t.output[len(t.output)-1]
}

//Empty Returns whether the queue is empty.
func (t *MyQueue) Empty() bool {
	isEmpty := false
	if len(t.input) == 0 && len(t.output) == 0 {
		isEmpty = true
	}
	return isEmpty
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
