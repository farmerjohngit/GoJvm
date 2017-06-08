package run

import (
	"math"
	"com/mogujie/wangzhi/jvmgo/run/heap"
)

type OperateStack struct {
	slots   []Slot
	top     uint32
	maxSize uint32
}

func NewOperateStack(maxSize uint32) *OperateStack {
	opStack := &OperateStack{
		maxSize: maxSize,
		top:     0,
		slots:   make([]Slot, maxSize)}
	return opStack
}

func (self *OperateStack) PushInt(val int32) {
	self.slots[self.top].num = val
	self.top++
}

func (self *OperateStack) PopInt() int32 {
	self.top--
	val := self.slots[self.top]

	return val.num
}

func (self *OperateStack) PushLong(val int64) {
	self.slots[self.top].num = int32(val)
	self.slots[self.top+1].num = int32(val >> 32)
	self.top += 2
}

func (self *OperateStack) PopLong() int64 {
	self.top--
	high := uint32(self.slots[self.top].num)
	self.top--
	low := uint32(self.slots[self.top].num)
	return int64(high)<<32 | int64(low)
}

func (self *OperateStack) PushFloat(val float32) {
	self.slots[self.top].num = int32(math.Float32bits(val))
	self.top++
}

func (self *OperateStack) PopFloat() float32 {
	self.top--
	val := self.slots[self.top]
	return math.Float32frombits(uint32(val.num))
}

func (self *OperateStack) PushDouble(val float64) {
	self.PushLong(int64(math.Float64bits(val)))
}

func (self *OperateStack) PopDouble() float64 {
	return math.Float64frombits(uint64(self.PopLong()))
}

func (self *OperateStack) PushRef(val *heap.Object) {
	self.slots[self.top].ref = val
	self.top ++
}

func (self *OperateStack) PopRef() *heap.Object {
	self.top--
	val := self.slots[self.top]
	return val.ref
}

func (self *OperateStack) PushSlot(slot Slot) {
	self.slots[self.top] = slot
	self.top++
}

func (self *OperateStack) PopSlot() Slot {
	self.top--
	val := self.slots[self.top]
	return val
}

func (self *OperateStack) TopSlot() Slot {
	return self.slots[self.top]
}
