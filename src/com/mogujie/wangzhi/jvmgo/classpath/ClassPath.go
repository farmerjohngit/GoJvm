package classpath

import (
	"path/filepath"
	"strings"
)

type ClassPath struct {
	CompositeEntry
}

func Parse(jrePath string, userClassParh string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootClassPath(jrePath)
	cp.parseUserClassPath(userClassParh)
	return cp
}

func (self *ClassPath) parseBootClassPath(jrePath string) {
	// jre/lib/*
	jreLibPath := filepath.Join(jrePath, "lib", "*")
	self.addEntry(newWildcardEntry(jreLibPath))

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jrePath, "lib", "ext", "*")
	self.addEntry(newWildcardEntry(jreExtPath))

}

func (self *ClassPath) parseUserClassPath(cpOption string) {
	self.addEntry(newEntry(cpOption))
}

func (self *ClassPath) ReadClass(className string) ([] byte, error) {
	//className.
	return self.readClass(strings.Replace(className, ".", pathSeparator, -1) + ".class")
}
