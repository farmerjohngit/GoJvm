package loads

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/run"
)

// Load reference from local variable
type ALOAD struct{ base.Index8Instruction }

func (self *ALOAD) Execute(frame *run.Frame) {
	_aload(frame, uint(self.Index))
}

type ALOAD_0 struct{ base.NoOpInstruction }

func (self *ALOAD_0) Execute(frame *run.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOpInstruction }

func (self *ALOAD_1) Execute(frame *run.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOpInstruction }

func (self *ALOAD_2) Execute(frame *run.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOpInstruction }

func (self *ALOAD_3) Execute(frame *run.Frame) {
	_aload(frame, 3)
}

func _aload(frame *run.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
