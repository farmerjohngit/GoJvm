package constant

import "com/mogujie/wangzhi/jvmgo/classfile"

type ConstantInfo interface {
	readInfo(reader *classfile.ClassReader)
}