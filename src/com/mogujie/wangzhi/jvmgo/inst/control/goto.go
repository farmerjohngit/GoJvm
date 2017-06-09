package control

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *run.Frame) {
	base.Branch(frame, self.Offset)
}
