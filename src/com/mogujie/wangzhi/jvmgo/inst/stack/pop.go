package stack

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type POP struct {
	base.NoOpInstruction
}

func (self *POP) Execute(frame *run.Frame) {
	frame.OpStack().PopSlot()
}

type POP2 struct {
	base.NoOpInstruction
}

func (self *POP2) Execute(frame *run.Frame) {
	frame.OpStack().PopSlot()
	frame.OpStack().PopSlot()
}
