package math

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Subtract double
type DSUB struct{ base.NoOpInstruction }

func (self *DSUB) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

// Subtract float
type FSUB struct{ base.NoOpInstruction }

func (self *FSUB) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// Subtract int
type ISUB struct{ base.NoOpInstruction }

func (self *ISUB) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

// Subtract long
type LSUB struct{ base.NoOpInstruction }

func (self *LSUB) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
