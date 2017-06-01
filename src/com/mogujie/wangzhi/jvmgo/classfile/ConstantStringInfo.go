package classfile


type ConstantStringInfo struct {
	index uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.index = reader.readUint16()
}

