package compare

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *run.Frame) {
	v1 := frame.OpStack().PopInt()
	if v1 == 0 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IFNE struct {
	base.BranchInstruction
}

func (self *IFNE) Execute(frame *run.Frame) {
	v1 := frame.OpStack().PopInt()
	if v1 != 0 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *run.Frame) {
	v1 := frame.OpStack().PopInt()
	if v1 < 0 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IFLE struct {
	base.BranchInstruction
}

func (self *IFLE) Execute(frame *run.Frame) {
	v1 := frame.OpStack().PopInt()
	if v1 <= 0 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IFGT struct {
	base.BranchInstruction
}

func (self *IFGT) Execute(frame *run.Frame) {
	v1 := frame.OpStack().PopInt()
	if v1 > 0 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IFGE struct {
	base.BranchInstruction
}

func (self *IFGE) Execute(frame *run.Frame) {
	v1 := frame.OpStack().PopInt()
	if v1 >= 0 {
		base.Branch(frame, int32(self.Offset))
	}
}
