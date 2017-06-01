package classfile

type ConstantLongInfo struct {
	val uint64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	self.val = reader.readUint64()
}
