package math

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Boolean XOR int
type IXOR struct{ base.NoOpInstruction }

func (self *IXOR) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long
type LXOR struct{ base.NoOpInstruction }

func (self *LXOR) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
