package control

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *run.Frame) {
	base.Branch(frame, int32(self.Offset))
}

type GOTOW struct {
	Offset int32
}

func (self *GOTOW) FetchOp(reader *base.ByteCodeReader) {
	self.Offset = reader.ReadInt32()
}

func (self *GOTOW) Execute(frame *run.Frame) {
	base.Branch(frame, self.Offset)
}
