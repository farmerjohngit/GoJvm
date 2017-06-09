package base

type ByteCodeReader struct {
	code []byte
	pc   int
}

func (self *ByteCodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *ByteCodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *ByteCodeReader) ReadUint8() uint8 {
	val := self.code[self.pc]
	self.pc++
	return val
}

func (self *ByteCodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

func (self *ByteCodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (self *ByteCodeReader) ReadInt32s(count int32) []int32 {
	ints := make([]int32, count)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

func (self *ByteCodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
