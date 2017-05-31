package classpath

import (
	"os"
	"strings"
	"com/mogujie/wangzhi/jvmgo/util"
)

const pathListSeparator = string(os.PathListSeparator)
const pathSeparator = string(os.PathSeparator)

type Entry interface {
	readClass(className string) ([] byte, error)
	toString() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if util.IsArchive(path) {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
