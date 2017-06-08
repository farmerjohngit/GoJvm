package stack

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type DUP struct {
	base.NoOpInstruction
}

func (self *DUP) Execute(frame *run.Frame) {
	slot := frame.OpStack().TopSlot()
	frame.OpStack().PushSlot(slot)
}

type DUP_X1 struct {
	base.NoOpInstruction
}

func (self *DUP_X1) Execute(frame *run.Frame) {
	//slot := frame.OpStack().TopSlot()
	val1 := frame.OpStack().PopSlot()
	val2 := frame.OpStack().PopSlot()
	frame.OpStack().PushSlot(val1)
	frame.OpStack().PushSlot(val2)
	frame.OpStack().PushSlot(val1)
}

type DUP_X2 struct {
	base.NoOpInstruction
}

func (self *DUP_X2) Execute(frame *run.Frame) {
	//slot := frame.OpStack().TopSlot()
	val1 := frame.OpStack().PopSlot()
	val2 := frame.OpStack().PopSlot()
	val3 := frame.OpStack().PopSlot()
	frame.OpStack().PushSlot(val1)
	frame.OpStack().PushSlot(val3)
	frame.OpStack().PushSlot(val2)
	frame.OpStack().PushSlot(val1)
}

type DUP2 struct {
	base.NoOpInstruction
}

func (self *DUP2) Execute(frame *run.Frame) {
	val1 := frame.OpStack().PopSlot()
	val2 := frame.OpStack().PopSlot()
	frame.OpStack().PushSlot(val2)
	frame.OpStack().PushSlot(val1)
	frame.OpStack().PushSlot(val2)
	frame.OpStack().PushSlot(val1)
}

type DUP2_X1 struct {
	base.NoOpInstruction
}

func (self *DUP2_X1) Execute(frame *run.Frame) {
	val1 := frame.OpStack().PopSlot()
	val2 := frame.OpStack().PopSlot()
	val3 := frame.OpStack().PopSlot()
	frame.OpStack().PushSlot(val2)
	frame.OpStack().PushSlot(val1)
	frame.OpStack().PushSlot(val3)
	frame.OpStack().PushSlot(val2)
	frame.OpStack().PushSlot(val1)
}

type DUP2_X2 struct {
	base.NoOpInstruction
}

func (self *DUP2_X2) Execute(frame *run.Frame) {
	val1 := frame.OpStack().PopSlot()
	val2 := frame.OpStack().PopSlot()
	val3 := frame.OpStack().PopSlot()
	val4 := frame.OpStack().PopSlot()
	frame.OpStack().PushSlot(val2)
	frame.OpStack().PushSlot(val1)
	frame.OpStack().PushSlot(val4)
	frame.OpStack().PushSlot(val3)
	frame.OpStack().PushSlot(val2)
	frame.OpStack().PushSlot(val1)

}
