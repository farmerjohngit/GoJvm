package conversions

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Convert long to double
type L2D struct{ base.NoOpInstruction }

func (self *L2D) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// Convert long to float
type L2F struct{ base.NoOpInstruction }

func (self *L2F) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// Convert long to int
type L2I struct{ base.NoOpInstruction }

func (self *L2I) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
