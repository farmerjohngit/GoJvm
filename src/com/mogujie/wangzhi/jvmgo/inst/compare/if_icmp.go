package compare

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *run.Frame) {
	v2 := frame.OpStack().PopInt()
	v1 := frame.OpStack().PopInt()
	if v1 == v2 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *run.Frame) {
	v2 := frame.OpStack().PopInt()
	v1 := frame.OpStack().PopInt()
	if v1 != v2 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (self *IF_ICMPLT) Execute(frame *run.Frame) {
	v2 := frame.OpStack().PopInt()
	v1 := frame.OpStack().PopInt()
	if v1 < v2 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPLE) Execute(frame *run.Frame) {
	v2 := frame.OpStack().PopInt()
	v1 := frame.OpStack().PopInt()
	if v1 <= v2 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *run.Frame) {
	v2 := frame.OpStack().PopInt()
	v1 := frame.OpStack().PopInt()
	if v1 > v2 {
		base.Branch(frame, int32(self.Offset))
	}
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *run.Frame) {
	v2 := frame.OpStack().PopInt()
	v1 := frame.OpStack().PopInt()
	if v1 >= v2 {
		base.Branch(frame, int32(self.Offset))
	}
}
