package conversions

import (
	"com/mogujie/wangzhi/jvmgo/run"
)

// Convert float to double
type F2D struct{ base.NoOpInstruction }

func (self *F2D) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// Convert float to int
type F2I struct{ base.NoOpInstruction }

func (self *F2I) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// Convert float to long
type F2L struct{ base.NoOpInstruction }

func (self *F2L) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
