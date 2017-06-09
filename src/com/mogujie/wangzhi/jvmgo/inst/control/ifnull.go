package control

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *run.Frame) {
	val := frame.OpStack().PopRef()
	if val == nil {
		base.Branch(frame, int32(self.Offset))
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNONNULL) Execute(frame *run.Frame) {
	ref := frame.OpStack().PopRef()
	if ref != nil {
		base.Branch(frame, int32(self.Offset))
	}
}
