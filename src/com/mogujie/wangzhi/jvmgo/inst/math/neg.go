package math

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Negate double
type DNEG struct{ base.NoOpInstruction }

func (self *DNEG) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ base.NoOpInstruction }

func (self *FNEG) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ base.NoOpInstruction }

func (self *INEG) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ base.NoOpInstruction }

func (self *LNEG) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
