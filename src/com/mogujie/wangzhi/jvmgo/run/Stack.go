package run

type Stack struct {
	top     *Frame
	maxSize int32
	size    int32
}

func NewStack(maxSize int32) *Stack {
	stack := &Stack{
		top:     nil,
		maxSize: maxSize,
		size:    0}
	return stack
}

func (self *Stack) pushFrame(frame *Frame) {
	frame.next = self.top
	self.top = frame
}

func (self *Stack) popFrame() *Frame {
	oldTop := self.top
	self.top = oldTop.next
	oldTop.next = nil
	return oldTop
}

func (self *Stack) currentFrame() *Frame {
	return self.top
}
