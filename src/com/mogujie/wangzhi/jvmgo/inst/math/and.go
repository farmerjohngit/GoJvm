package math

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Boolean AND int
type IAND struct{ base.NoOpInstruction }

func (self *IAND) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type LAND struct{ base.NoOpInstruction }

func (self *LAND) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
