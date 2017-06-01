package classfile

import (
	"fmt"
	//"reflect"
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
	constantInfo []ConstantInfo
}

func NewConstantPool(count uint16) *ConstantPool {
	pool := &ConstantPool{
		constantInfo: make([]ConstantInfo, count)}

	return pool
}

func (self *ConstantPool) getDescriptor(index uint16) string {
	return self.getUtf8(self.constantInfo[index].(*ConstantNameAndTypeInfo).descriptorIndex)
}

func (self *ConstantPool) getNameAndType(index uint16) (name, _type string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name = self.getUtf8(ntInfo.nameIndex)
	_type = self.getUtf8(ntInfo.descriptorIndex)
	return
}

func (self *ConstantPool) getClassName(index uint16) string {
	return self.getUtf8(self.constantInfo[index].(*ConstantClassInfo).nameIndex)
}

func (self *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	return self.constantInfo[index]
}

func (self *ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (self *ConstantPool) read(reader *ClassReader) {
	for i := 1; i < len(self.constantInfo); i++ {
		tag := reader.ReadUint8()
		cInfo := self.newConstantInfo(tag)
		cInfo.readInfo(reader)
		self.constantInfo[i] = cInfo
	}
}

func (self *ConstantPool) printConstantPool() {

	for i := 1; i < len(self.constantInfo); i++ {
		info := self.constantInfo[i]
		//println("type: ", reflect.TypeOf(info).String())

		switch info.(type) {
		case *ConstantUtf8Info:
			println("utf8:", info.(*ConstantUtf8Info).str)
		case *ConstantClassInfo:
			println("class:", self.getUtf8(info.(*ConstantClassInfo).nameIndex))
		case *ConstantDoubleInfo:
			println("double:", info.(*ConstantDoubleInfo).val)
		case *ConstantFloatInfo:
			println("float:", info.(*ConstantFloatInfo).val)
		case *ConstantIntegerInfo:
			println("integer:", info.(*ConstantIntegerInfo).val)
		case *ConstantLongInfo:
			println("long:", info.(*ConstantLongInfo).val)
		case *ConstantMemberRefInfo:
			name, typeStr := self.getNameAndType(info.(*ConstantMemberRefInfo).nameAndTypeIndex)
			println("member_ref:", self.getClassName(info.(*ConstantMemberRefInfo).classIndex), name, typeStr)
		case *ConstantMethodRefInfo:
			name, typeStr := self.getNameAndType(info.(*ConstantMethodRefInfo).nameAndTypeIndex)
			println("method_ref:", self.getClassName(info.(*ConstantMethodRefInfo).classIndex), name, typeStr)
		case *ConstantFieldRefInfo:
			name, typeStr := self.getNameAndType(info.(*ConstantFieldRefInfo).nameAndTypeIndex)
			println("field_ref:", self.getClassName(info.(*ConstantFieldRefInfo).classIndex), name, typeStr)
		case *ConstantNameAndTypeInfo:
			name, typeStr := self.getNameAndType(uint16(i))
			println("name_and_type:", name, typeStr)
		case *ConstantStringInfo:
			println("string:", self.getUtf8(info.(*ConstantStringInfo).index))
		default:
			println("other")
		}
	}
}

func (self *ConstantPool) newConstantInfo(tag uint8) ConstantInfo {
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
		return &ConstantStringInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default: // todo
		fmt.Errorf("BAD constant pool tag: %v", tag)
		return nil
	}
}
