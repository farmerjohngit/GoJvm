package jvmgo

import (
	"com/mogujie/wangzhi/jvmgo/cmd"
	"com/mogujie/wangzhi/jvmgo/classpath"
	"com/mogujie/wangzhi/jvmgo/classfile"
	"fmt"
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
	println("bytes:", bytes)
	println("magic number:", fmt.Sprintf("%X", cf.Magic()))
	println("MinorVersion:", fmt.Sprintf("%X", cf.MinorVersion()))
	println("MajorVersion:", fmt.Sprintf("%X", cf.MajorVersion()))

}
