package math

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
	"math"
)

// Add float
type IREM struct{ base.NoOpInstruction }

func (self *IREM) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2==0{
	panic("java.lang.ArithmeticException: / by zero")
	}

	stack.PushInt(v1 % v2)
}

// Remainder double
type DREM struct{ base.NoOpInstruction }

func (self *DREM) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2) // todo
	stack.PushDouble(result)
}

// Remainder float
type FREM struct{ base.NoOpInstruction }

func (self *FREM) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

// Remainder long
type LREM struct{ base.NoOpInstruction }

func (self *LREM) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 % v2
		stack.PushLong(result)
	}
}
