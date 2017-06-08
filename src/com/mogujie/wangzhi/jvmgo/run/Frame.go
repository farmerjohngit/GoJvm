package run

type Frame struct {
	next      *Frame
	opStack   *OperateStack
	localVars *LocalVariables
}

func NewFrame(maxOpStack uint32, maxLocalVariable uint32) *Frame {
	frame := &Frame{
		next:      nil,
		opStack:   NewOperateStack(maxOpStack),
		localVars: NewLocalVariable(maxLocalVariable)}

	return frame
}

func (self *Frame) LocalVars() *LocalVariables {
	return self.localVars
}

func (self *Frame) OpStack() *OperateStack {
	return self.opStack
}
