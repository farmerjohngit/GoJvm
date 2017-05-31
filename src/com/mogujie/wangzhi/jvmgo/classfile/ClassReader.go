package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func newClassReader(data []byte) *ClassReader {
	return &ClassReader{data}
}

// u1
func (self *ClassReader) ReadUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// u2
func (self *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// u4
func (self *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) ReadUint16Arr() []uint16 {
	n := self.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.ReadUint16()
	}
	return s
}

func (self *ClassReader) ReadBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}
