package stores

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Store long into local variable
type LSTORE struct{ base.Index8Instruction }

func (self *LSTORE) Execute(frame *run.Frame) {
	_lstore(frame, uint(self.Index))
}

type LSTORE_0 struct{ base.NoOpInstruction }

func (self *LSTORE_0) Execute(frame *run.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ base.NoOpInstruction }

func (self *LSTORE_1) Execute(frame *run.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ base.NoOpInstruction }

func (self *LSTORE_2) Execute(frame *run.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ base.NoOpInstruction }

func (self *LSTORE_3) Execute(frame *run.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *run.Frame, index uint) {
	val := frame.OpStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
