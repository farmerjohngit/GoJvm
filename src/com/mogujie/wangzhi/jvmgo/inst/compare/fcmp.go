package compare

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type FCMPG struct {
	base.NoOpInstruction
}

func (self *FCMPG) Execute(frame *run.Frame) {
	self._fcmp(frame, true)
}

func (self *FCMPG) _fcmp(frame *run.Frame, flag bool) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
