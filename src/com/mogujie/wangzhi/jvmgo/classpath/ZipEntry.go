package classpath

import (
	"path/filepath"
)

type ZipEntry struct {
	zipPath string
}

//todo
func (self *ZipEntry) readClass(className string) ([] byte, error) {
	file := filepath.Join(self.zipPath, className)

	println(file)
	return nil, nil
}

func (self *ZipEntry) toString() string {
	return self.zipPath
}

func newZipEntry(path string) *ZipEntry {
	zipPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{zipPath}
}
