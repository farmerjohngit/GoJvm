package main

import (
	"com/mogujie/wangzhi/jvmgo"
	"com/mogujie/wangzhi/jvmgo/cmd"
)

func main() {

	arg := []string{"java", "-cp", "/Users/farmerjohn/IdeaProjects/pandora-gradle/build/classes/main", "-XjavaHome",
			"/Library/Java/JavaVirtualMachines/jdk1.8.0_111.jdk/Contents/Home", "-verbose","com.wangzhi.jvmgo.Main"}
	cmd, err := cmd.ParseCommand(arg)
	if err != nil {
		panic(err)
	}
	jvmgo.StartJvm(cmd)
}
