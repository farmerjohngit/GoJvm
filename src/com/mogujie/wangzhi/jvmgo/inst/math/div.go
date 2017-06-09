package math

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Divide double
type DDIV struct{ base.NoOpInstruction }

func (self *DDIV) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// Divide float
type FDIV struct{ base.NoOpInstruction }

func (self *FDIV) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// Divide int
type IDIV struct{ base.NoOpInstruction }

func (self *IDIV) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 / v2
		stack.PushInt(result)
	}
}

// Divide long
type LDIV struct{ base.NoOpInstruction }

func (self *LDIV) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 / v2
		stack.PushLong(result)
	}
}
