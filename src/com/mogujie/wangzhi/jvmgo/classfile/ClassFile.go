package classfile

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
	pool         *ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	AttributeTable
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readMagic(reader)
	self.readVersion(reader)
	self.readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16Arr()
	self.fields = self.readMembers(reader)
	self.methods = self.readMembers(reader)
	self.attributes = readAttributes(reader, self.pool)
}

func (self *ClassFile) readMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
}

func (self *ClassFile) readVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
}

func (self *ClassFile) readConstantPool(reader *ClassReader) {
	self.poolCount = reader.readUint16()
	self.pool = NewConstantPool(self.poolCount)
	self.pool.read(reader)
	self.pool.printConstantPool()
}

// read field or method table
func (self *ClassFile) readMembers(reader *ClassReader) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = &MemberInfo{cp: self.pool}
		members[i].read(reader)
	}
	return members
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
