package stores


import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Store double into local variable
type DSTORE struct{ base.Index8Instruction }

func (self *DSTORE) Execute(frame *run.Frame) {
	_dstore(frame, uint(self.Index))
}

type DSTORE_0 struct{ base.NoOpInstruction }

func (self *DSTORE_0) Execute(frame *run.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOpInstruction }

func (self *DSTORE_1) Execute(frame *run.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOpInstruction }

func (self *DSTORE_2) Execute(frame *run.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOpInstruction }

func (self *DSTORE_3) Execute(frame *run.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *run.Frame, index uint) {
	val := frame.OpStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
