package constant

import (
	"fmt"
	"com/mogujie/wangzhi/jvmgo/classfile"
)

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantPool struct {
	count        uint16
	constantInfo []ConstantInfo
}

func NewConstantPool(count uint16) *ConstantPool {
	pool := &ConstantPool{
		count:        count,
		constantInfo: make([]ConstantInfo, count)}

	return pool
}

func (self *ConstantPool) read(reader *classfile.ClassReader) {
	tag := reader.ReadUint8()
	constantInfo := self.getConstantInfo(tag)
	constantInfo.readInfo(reader)
}

func (self *ConstantPool) getConstantInfo(tag uint8) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: self}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: self}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: self}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: self}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: self}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{cp: self}
	default: // todo
		fmt.Errorf("BAD constant pool tag: %v", tag)
		return nil
	}
}
