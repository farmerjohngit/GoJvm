package constant

import (
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
)

type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOp(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(int32(self.val))
}

type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOp(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *run.Frame) {
	frame.OpStack().PushInt(int32(self.val))
}
