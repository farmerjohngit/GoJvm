package math

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Boolean OR int
type IOR struct{ base.NoOpInstruction }

func (self *IOR) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ base.NoOpInstruction }

func (self *LOR) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
