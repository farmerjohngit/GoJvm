package stores

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Store int into local variable
type ISTORE struct{ base.Index8Instruction }

func (self *ISTORE) Execute(frame *run.Frame) {
	_istore(frame, uint(self.Index))
}

type ISTORE_0 struct{ base.NoOpInstruction }

func (self *ISTORE_0) Execute(frame *run.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct{ base.NoOpInstruction }

func (self *ISTORE_1) Execute(frame *run.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct{ base.NoOpInstruction }

func (self *ISTORE_2) Execute(frame *run.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct{ base.NoOpInstruction }

func (self *ISTORE_3) Execute(frame *run.Frame) {
	_istore(frame, 3)
}

func _istore(frame *run.Frame, index uint) {
	val := frame.OpStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
