package classfile

type MemberInfo struct {
	cp              *ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	AttributeTable
}

func (self *MemberInfo) read(reader *ClassReader) {
	self.accessFlags = reader.readUint16()
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
	self.attributes = readAttributes(reader, self.cp)
}
