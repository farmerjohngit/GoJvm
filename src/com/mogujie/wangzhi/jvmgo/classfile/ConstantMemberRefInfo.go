package classfile


/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMemberRefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}


func (self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

type MemberRefInfo interface {

}