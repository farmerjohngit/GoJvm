package classfile

import "com/mogujie/wangzhi/jvmgo/classfile/constant"

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	poolCount    uint16
	pool         *constant.ConstantPool
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readMagic(reader)
	self.readVersion(reader)
}

func (self *ClassFile) readMagic(reader *ClassReader) {
	self.magic = reader.ReadUint32()
}

func (self *ClassFile) readVersion(reader *ClassReader) {
	self.minorVersion = reader.ReadUint16()
	self.majorVersion = reader.ReadUint16()
}

func (self *ClassFile) readConstantPool(reader *ClassReader) {
	self.poolCount = reader.ReadUint16()
	self.pool = constant.NewConstantPool(self.poolCount)
}

/********************The Get Method********************/

func (self *ClassFile) Magic() uint32 {
	return self.magic
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
