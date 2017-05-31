package constant

import "com/mogujie/wangzhi/jvmgo/classfile"

type ConstantStringInfo struct {
}

func (*ConstantStringInfo) readInfo(reader *classfile.ClassReader) {
	panic("implement me")
}
