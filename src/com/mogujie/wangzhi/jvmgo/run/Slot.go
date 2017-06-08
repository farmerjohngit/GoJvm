package run

import "com/mogujie/wangzhi/jvmgo/run/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
