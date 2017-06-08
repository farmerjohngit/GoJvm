package loads

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Load long from local variable
type LLOAD struct{ base.Index8Instruction }

func (self *LLOAD) Execute(frame *run.Frame) {
	_lload(frame, uint(self.Index))
}

type LLOAD_0 struct{ base.NoOpInstruction }

func (self *LLOAD_0) Execute(frame *run.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOpInstruction }

func (self *LLOAD_1) Execute(frame *run.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOpInstruction }

func (self *LLOAD_2) Execute(frame *run.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOpInstruction }

func (self *LLOAD_3) Execute(frame *run.Frame) {
	_lload(frame, 3)
}

func _lload(frame *run.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OpStack().PushLong(val)
}
