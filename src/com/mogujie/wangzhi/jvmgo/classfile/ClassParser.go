package classfile

type ClassParser struct {
}

func Parse(bytes []byte) (cf *ClassFile, err error) {
	reader := &ClassReader{bytes}
	classFile := &ClassFile{}
	classFile.read(reader)
	return classFile, nil
}


