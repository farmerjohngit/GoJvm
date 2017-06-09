package conversions

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Convert double to float
type D2F struct{ base.NoOpInstruction }

func (self *D2F) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOpInstruction }

func (self *D2I) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOpInstruction }

func (self *D2L) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
