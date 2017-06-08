package stack

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type SWAP struct {
	base.NoOpInstruction
}

func (self *SWAP) Execute(frame *run.Frame) {
	val1 := frame.OpStack().PopSlot()
	val2 := frame.OpStack().PopSlot()
	frame.OpStack().PushSlot(val1)
	frame.OpStack().PushSlot(val2)
}
