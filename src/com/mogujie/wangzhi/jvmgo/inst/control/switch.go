package control

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITCH) FetchOp(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	self.jumpOffsets = reader.ReadInt32s(self.high - self.low + 1)
}

func (self *TABLE_SWITCH) Execute(frame *run.Frame) {
	key := frame.OpStack().PopInt()
	if key >= self.low && key <= self.high {
		base.Branch(frame, self.jumpOffsets[key-self.low])
	} else {
		base.Branch(frame, self.defaultOffset)
	}
}

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOp(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *run.Frame) {
	key := frame.OpStack().PopInt()

	for i := int32(0); i < self.npairs*2; i += 2 {
		if key == self.matchOffsets[i] {
			base.Branch(frame, self.matchOffsets[i+1])
			return
		}
	}
	base.Branch(frame, self.defaultOffset)
}
