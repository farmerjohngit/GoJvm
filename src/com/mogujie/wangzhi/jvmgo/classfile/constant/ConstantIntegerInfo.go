package constant

import (
	"com/mogujie/wangzhi/jvmgo/classfile"
)

type ConstantIntegerInfo struct {
}

func (*ConstantIntegerInfo) readInfo(reader *classfile.ClassReader) {
	panic("implement me")
}
