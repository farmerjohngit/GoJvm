package compare

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *run.Frame) {
	if _acmp(frame) {
		base.Branch(frame, int32(self.Offset))
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *run.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, int32(self.Offset))
	}
}

func _acmp(frame *run.Frame) bool {
	stack := frame.OpStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
