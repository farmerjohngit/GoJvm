package classfile

type ConstantIntegerInfo struct {
	val uint32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	self.val = reader.readUint32()
}
