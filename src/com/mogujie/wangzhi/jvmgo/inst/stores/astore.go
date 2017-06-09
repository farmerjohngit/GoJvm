package stores

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Store reference into local variable
type ASTORE struct{ base.Index8Instruction }

func (self *ASTORE) Execute(frame *run.Frame) {
	_astore(frame, uint(self.Index))
}

type ASTORE_0 struct{ base.NoOpInstruction }

func (self *ASTORE_0) Execute(frame *run.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOpInstruction }

func (self *ASTORE_1) Execute(frame *run.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOpInstruction }

func (self *ASTORE_2) Execute(frame *run.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOpInstruction }

func (self *ASTORE_3) Execute(frame *run.Frame) {
	_astore(frame, 3)
}

func _astore(frame *run.Frame, index uint) {
	ref := frame.OpStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
