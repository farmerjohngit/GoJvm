package main


type test struct {
	num int32

}

func main() {
	objs:=make([]test,3)

	println(objs[0].num)
}
