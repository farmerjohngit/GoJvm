package run

const default_max_stack = 256

type Thread struct {
	stack *Stack
	pc    int
}

func NewThread() *Thread {
	thread := &Thread{
		pc:    -1,
		stack: NewStack(default_max_stack)}
	return thread
}

func NewThreadWithMaxStack(maxStack int32) *Thread {
	thread := &Thread{
		pc:    -1,
		stack: NewStack(maxStack)}
	return thread
}

func (self *Thread) pushFrame(frame *Frame) {
	self.stack.pushFrame(frame)
}

func (self *Thread) popFrame() *Frame {
	return self.stack.popFrame()
}
