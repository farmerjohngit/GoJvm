package conversions

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

// Convert int to byte
type I2B struct{ base.NoOpInstruction }

func (self *I2B) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

// Convert int to char
type I2C struct{ base.NoOpInstruction }

func (self *I2C) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}

// Convert int to short
type I2S struct{ base.NoOpInstruction }

func (self *I2S) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

// Convert int to long
type I2L struct{ base.NoOpInstruction }

func (self *I2L) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

// Convert int to float
type I2F struct{ base.NoOpInstruction }

func (self *I2F) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

// Convert int to double
type I2D struct{ base.NoOpInstruction }

func (self *I2D) Execute(frame *run.Frame) {
	stack := frame.OpStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
