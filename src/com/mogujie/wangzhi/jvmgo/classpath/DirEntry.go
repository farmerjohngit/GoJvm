package classpath

import (
	"path/filepath"
	"io/ioutil"
	"os"
)

type DirEntry struct {
	dirPath string
}

func (dirEntry *DirEntry) readClass(className string) ([] byte, error) {
	file := filepath.Join(dirEntry.dirPath, className)
	return ioutil.ReadFile(file)
}

func (self *DirEntry) toString() string {
	return self.dirPath
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	file, err := os.Stat(absDir)
	if !file.IsDir() {
		panic(path + " is not dir!")
	}

	return &DirEntry{absDir}
}
