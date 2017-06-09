package jvmgo

import (
	"com/mogujie/wangzhi/jvmgo/classfile"
	"com/mogujie/wangzhi/jvmgo/run"
	"com/mogujie/wangzhi/jvmgo/inst/base"
	"com/mogujie/wangzhi/jvmgo/inst"
	"fmt"
)

func interpret(methodInfo *classfile.MemberInfo) {
	code := methodInfo.CodeAttribute()
	thread := run.NewThread()
	frame := run.NewFrame(thread, code.MaxStack(), code.MaxLocals())
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, code.Code())
}

func catchErr(frame *run.Frame) {
	if r := recover(); r != nil {
		println(fmt.Sprintf("LocalVars:%v\n", frame.LocalVars()))
		println(fmt.Sprintf("OpStack:%v\n", frame.OpStack()))
		panic(r)
	}
}

func loop(thread *run.Thread, code []byte) {
	reader := &base.ByteCodeReader{}

	frame := thread.PopFrame()
	for {
		pc := frame.NextPc()
		thread.SetPC(pc)
		reader.Reset(code, pc)
		opcode := reader.ReadUint8()
		inst := inst.NewInstruction(opcode)

		println("opcode: ", opcode)
		inst.FetchOp(reader)
		frame.SetNextPc(reader.PC())
		inst.Execute(frame)
	}

}
