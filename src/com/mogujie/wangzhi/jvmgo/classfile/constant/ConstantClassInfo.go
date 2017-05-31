package constant

import "com/mogujie/wangzhi/jvmgo/classfile"

type ConstantClassInfo struct {
}

func (*ConstantClassInfo) readInfo(reader *classfile.ClassReader) {
	panic("implement me")
}
