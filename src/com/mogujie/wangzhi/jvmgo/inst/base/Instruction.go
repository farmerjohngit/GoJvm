package base

import "com/mogujie/wangzhi/jvmgo/run"

type Instruction interface {
	FetchOp(reader *ByteCodeReader)
	Execute(frame *run.Frame)
}

type NoOpInstruction struct {
}

func (self *NoOpInstruction) FetchOp(reader *ByteCodeReader) {
	//do nothing
}

type BranchInstruction struct {
	Offset int16
}

func (self *BranchInstruction) FetchOp(reader *ByteCodeReader) {
	self.Offset = reader.ReadInt16()
}

type Index8Instruction struct {
	//because  there have a inst 'wide', so we can not use a uint8 represent index
	Index uint
}

func (self *Index8Instruction) FetchOp(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOp(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint16())
}
