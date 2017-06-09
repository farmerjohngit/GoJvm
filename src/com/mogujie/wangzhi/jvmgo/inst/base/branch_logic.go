package base

import "com/mogujie/wangzhi/jvmgo/run"

func Branch(frame *run.Frame, offset int32) {
	frame.SetNextPc(frame.Thread().PC() + offset)
}
