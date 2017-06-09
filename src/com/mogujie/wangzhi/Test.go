package main

import (
	"fmt"
	"encoding/binary"
	"bytes"
)

type test struct {
	num int32
}

func main() {
	//bs := make([]byte, 4)
	//binary.LittleEndian.PutUint32(bs, 65531)
	//fmt.Println(bs)
	buf := new(bytes.Buffer)
	var num uint16 = 65531
	err := binary.Write(buf, binary.BigEndian, int32(num))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% b", buf.Bytes())
	//var i uint16
	//i = 65531
	//fmt.Printf("%b\n", i)
	//fmt.Printf("%b\n", int16(i))
	//fmt.Printf("%b\n", int(int16(i)))
}
