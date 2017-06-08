package loads

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Load double from local variable
type DLOAD struct{ base.Index8Instruction }

func (self *DLOAD) Execute(frame *run.Frame) {
	_dload(frame, uint(self.Index))
}

type DLOAD_0 struct{ base.NoOpInstruction }

func (self *DLOAD_0) Execute(frame *run.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOpInstruction }

func (self *DLOAD_1) Execute(frame *run.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOpInstruction }

func (self *DLOAD_2) Execute(frame *run.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOpInstruction }

func (self *DLOAD_3) Execute(frame *run.Frame) {
	_dload(frame, 3)
}

func _dload(frame *run.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OpStack().PushDouble(val)
}
