package compare

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

type DCMPG struct{ base.NoOpInstruction }

func (self *DCMPG) Execute(frame *run.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ base.NoOpInstruction }

func (self *DCMPL) Execute(frame *run.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *run.Frame, gFlag bool) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
