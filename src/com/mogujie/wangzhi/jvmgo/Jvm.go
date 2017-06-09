package jvmgo

import (
	"com/mogujie/wangzhi/jvmgo/cmd"

	"com/mogujie/wangzhi/jvmgo/run"
	"fmt"
	"com/mogujie/wangzhi/jvmgo/run/heap"
	"com/mogujie/wangzhi/jvmgo/classpath"
	"com/mogujie/wangzhi/jvmgo/classfile"
)

func StartJvm(cmd *cmd.Command) {
	cp := classpath.Parse(cmd.Options().JavaHome(), cmd.Options().Classpath())
	bytes, err := cp.ReadClass("com.wangzhi.jvmgo.Main")
	if err != nil {
		panic(err)
	}
	println("------------------------")
	cf, err := classfile.Parse(bytes)
	if err != nil {
		panic(err)
	}
	mtds := cf.MethodInfo()

	for i := range mtds {
		mtd := mtds[i]
		println("mtd:Name", mtd.Name())
		println("mtd:Descriptor", mtd.Descriptor())

		if mtd.Name() == "main" && mtd.Descriptor() == "([Ljava/lang/String;)V" {
			interpret(mtd)
		}

	}

	//println("bytes:", bytes)
	//println("magic number:", fmt.Sprintf("%X", cf.Magic()))
	//println("MinorVersion:", fmt.Sprintf("%X", cf.MinorVersion()))
	//println("MajorVersion:", fmt.Sprintf("%X", cf.MajorVersion()))

	//frame := run.NewFrame(32, 32)

	//testLocalVars(frame.LocalVars())
	//testOpStack(frame.OpStack())

}

func testLocalVars(vars *run.LocalVariables) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.141)
	vars.SetDouble(7, 3.14159265358979)
	vars.SetRef(9, nil)

	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

//
func testOpStack(stack *run.OperateStack) {
	stack.PushInt(100)
	stack.PushInt(-100)
	stack.PushLong(100)
	stack.PushLong(-100)
	stack.PushFloat(3.141)
	stack.PushDouble(3.14159265358979)
	stack.PushDouble(-3.14159265358979)
	stack.PushRef(&heap.Object{})
	fmt.Printf("%d\n", stack.PopRef())
	fmt.Printf("%f\n", stack.PopDouble())
	fmt.Printf("%f\n", stack.PopDouble())
	fmt.Printf("%f\n", stack.PopFloat())
	fmt.Printf("%d\n", stack.PopLong())
	fmt.Printf("%d\n", stack.PopLong())
	fmt.Printf("%d\n", stack.PopInt())
	fmt.Printf("%d\n", stack.PopInt())

}
