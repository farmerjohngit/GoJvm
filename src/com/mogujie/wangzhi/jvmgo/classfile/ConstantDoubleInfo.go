package classfile

import (
	"math"
)

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	self.val = math.Float64frombits(reader.readUint64())
}
