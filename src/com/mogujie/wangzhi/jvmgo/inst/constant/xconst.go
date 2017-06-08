package constant

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type ACONST_NULL struct {
	base.NoOpInstruction
}

func (self *ACONST_NULL) Execute(frame *run.Frame) {
	frame.OpStack().PushRef(nil)
}

type DCONST_0 struct {
	base.NoOpInstruction
}

func (self *DCONST_0) Execute(frame *run.Frame) {
	frame.OpStack().PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOpInstruction }

func (self *DCONST_1) Execute(frame *run.Frame) {
	frame.OpStack().PushDouble(1.0)
}

// Push float
type FCONST_0 struct{ base.NoOpInstruction }

func (self *FCONST_0) Execute(frame *run.Frame) {
	frame.OpStack().PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOpInstruction }

func (self *FCONST_1) Execute(frame *run.Frame) {
	frame.OpStack().PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOpInstruction }

func (self *FCONST_2) Execute(frame *run.Frame) {
	frame.OpStack().PushFloat(2.0)
}

type ICONST_M1 struct {
	base.NoOpInstruction
}

func (self *ICONST_M1) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(-1)
}

type ICONST_0 struct{ base.NoOpInstruction }

func (self *ICONST_0) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(0)
}

type ICONST_1 struct{ base.NoOpInstruction }

func (self *ICONST_1) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(1)
}

type ICONST_2 struct{ base.NoOpInstruction }

func (self *ICONST_2) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(2)
}

type ICONST_3 struct{ base.NoOpInstruction }

func (self *ICONST_3) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOpInstruction }

func (self *ICONST_4) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(4)
}

type ICONST_5 struct{ base.NoOpInstruction }

func (self *ICONST_5) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(5)
}

// Push long constant
type LCONST_0 struct{ base.NoOpInstruction }

func (self *LCONST_0) Execute(frame *run.Frame) {
	frame.OpStack().PushLong(0)
}

type LCONST_1 struct{ base.NoOpInstruction }

func (self *LCONST_1) Execute(frame *run.Frame) {
	frame.OpStack().PushLong(1)
}
