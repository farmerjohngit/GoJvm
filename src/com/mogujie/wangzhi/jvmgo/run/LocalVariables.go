package run

import (
	"com/mogujie/wangzhi/jvmgo/run/heap"
	"math"
)

type LocalVariables struct {
	slots   []Slot
	maxSize uint32
}

func NewLocalVariable(maxSize uint32) *LocalVariables {
	opStack := &LocalVariables{
		maxSize: maxSize,
		slots:   make([]Slot, maxSize)}
	return opStack
}

func (self *LocalVariables) GetRef(index uint) *heap.Object {
	return self.slots[index].ref
}
func (self *LocalVariables) SetRef(index uint, ref *heap.Object) {
	self.slots[index].ref = ref
}

func (self *LocalVariables) GetBoolean(index uint) bool {
	return self.GetInt(index) == 1
}

func (self *LocalVariables) GetInt(index uint) int32 {
	return self.slots[index].num
}
func (self *LocalVariables) SetInt(index uint, val int32) {
	self.slots[index].num = val
}

func (self *LocalVariables) GetLong(index uint) int64 {
	low := uint32(self.slots[index].num)
	high := uint32(self.slots[index+1].num)

	return int64(high)<<32 | int64(low)
}

func (self *LocalVariables) SetLong(index uint, val int64) {
	self.slots[index].num = int32(val)
	self.slots[index+1].num = int32(val >> 32)
}

func (self *LocalVariables) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(self.GetInt(index)))
}

func (self *LocalVariables) SetFloat(index uint, val float32) {
	self.SetInt(index, int32(math.Float32bits(val)))
}

func (self *LocalVariables) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(self.GetLong(index)))
}
func (self *LocalVariables) SetDouble(index uint, val float64) {
	self.SetLong(index, int64(math.Float64bits(val)))
}

func (self *LocalVariables) Get(index uint) *heap.Object {
	return self.slots[index].ref
}

func (self *LocalVariables) Set(index uint, any *heap.Object) {
	self.slots[index].ref = any
}
