package run

type Frame struct {
	next      *Frame
	opStack   *OperateStack
	localVars *LocalVariables
	thread    *Thread
	nextPC    int32
}

func NewFrame(maxOpStack uint32, maxLocalVariable uint32) *Frame {
	frame := &Frame{
		next:      nil,
		opStack:   NewOperateStack(maxOpStack),
		localVars: NewLocalVariable(maxLocalVariable),
		thread:    NewThread()}

	return frame
}

func (self *Frame) LocalVars() *LocalVariables {
	return self.localVars
}

func (self *Frame) OpStack() *OperateStack {
	return self.opStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPc(nextPC int32) { self.nextPC = nextPC }
