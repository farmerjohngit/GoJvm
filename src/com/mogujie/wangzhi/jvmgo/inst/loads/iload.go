package loads

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type ILOAD_0 struct {
	base.NoOpInstruction
}

func (self *ILOAD_0) Execute(frame *run.Frame) {
	_iload(frame,0)
}

type ILOAD_1 struct {
	base.NoOpInstruction
}

func (self *ILOAD_1) Execute(frame *run.Frame) {
	_iload(frame,1)
}

type ILOAD_2 struct {
	base.NoOpInstruction
}

func (self *ILOAD_2) Execute(frame *run.Frame) {
	_iload(frame,2)
}

type ILOAD_3 struct {
	base.NoOpInstruction
}

func (self *ILOAD_3) Execute(frame *run.Frame) {
	_iload(frame,3)
}


type ILOAD_X struct {
	base.Index8Instruction
}

func (self *ILOAD_X) Execute(frame *run.Frame) {
	_iload(frame,uint(self.Index))
}


func _iload(frame *run.Frame,index uint){
	frame.OpStack().PushInt(frame.LocalVars().GetInt(index))
}

