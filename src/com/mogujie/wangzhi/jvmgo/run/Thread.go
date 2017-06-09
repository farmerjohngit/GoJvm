package run

const default_max_stack = 256

type Thread struct {
	stack *Stack
	pc    int32
}

func NewThread() *Thread {
	thread := &Thread{
		pc:    0,
		stack: NewStack(default_max_stack)}
	return thread
}

func NewThreadWithMaxStack(maxStack int32) *Thread {
	thread := &Thread{
		pc:    -1,
		stack: NewStack(maxStack)}
	return thread
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.pushFrame(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.popFrame()
}

func (self *Thread) SetPC(pc int32) {
	self.pc = pc
}

func (self *Thread) PC() int32 {
	return self.pc
}
