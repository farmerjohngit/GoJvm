package loads

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Load float from local variable
type FLOAD struct{ base.Index8Instruction }

func (self *FLOAD) Execute(frame *run.Frame) {
	_fload(frame, uint(self.Index))
}

type FLOAD_0 struct{ base.NoOpInstruction }

func (self *FLOAD_0) Execute(frame *run.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOpInstruction }

func (self *FLOAD_1) Execute(frame *run.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOpInstruction }

func (self *FLOAD_2) Execute(frame *run.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOpInstruction }

func (self *FLOAD_3) Execute(frame *run.Frame) {
	_fload(frame, 3)
}

func _fload(frame *run.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OpStack().PushFloat(val)
}
