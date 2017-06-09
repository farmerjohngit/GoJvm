package math

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type IADD struct {
	base.NoOpInstruction
}

func (slef *IADD) Execute(frame *run.Frame) {

	val1 := frame.OpStack().PopInt()
	val2 := frame.OpStack().PopInt()
	frame.OpStack().PushInt(val1 + val2)
}

// Add double
type DADD struct{ base.NoOpInstruction }

func (self *DADD) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// Add float
type FADD struct{ base.NoOpInstruction }

func (self *FADD) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// Add long
type LADD struct{ base.NoOpInstruction }

func (self *LADD) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
