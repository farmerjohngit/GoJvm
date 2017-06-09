package stores

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Store float into local variable
type FSTORE struct{ base.Index8Instruction }

func (self *FSTORE) Execute(frame *run.Frame) {
	_fstore(frame, uint(self.Index))
}

type FSTORE_0 struct{ base.NoOpInstruction }

func (self *FSTORE_0) Execute(frame *run.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOpInstruction }

func (self *FSTORE_1) Execute(frame *run.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOpInstruction }

func (self *FSTORE_2) Execute(frame *run.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOpInstruction }

func (self *FSTORE_3) Execute(frame *run.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *run.Frame, index uint) {
	val := frame.OpStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
