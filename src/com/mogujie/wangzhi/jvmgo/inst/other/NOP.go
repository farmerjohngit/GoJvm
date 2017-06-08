package other

import (
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/run"
)

type NOP struct {
	base.NoOpInstruction
}

func (self *NOP) Execute(frame *run.Frame) {

}
