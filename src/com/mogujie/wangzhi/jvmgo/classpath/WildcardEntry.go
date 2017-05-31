package classpath

import (
	"path/filepath"
	"os"
	"com/mogujie/wangzhi/jvmgo/util"
)

type WildcardEntry struct {
	CompositeEntry
}

func newWildcardEntry(path string) *WildcardEntry {
	baseDir := path[:len(path)-1] // remove *
	entry := &WildcardEntry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if util.IsArchive(path) {
			jarEntry := newZipEntry(path)
			entry.addEntry(jarEntry)
		}

		return nil
	}

	filepath.Walk(baseDir, walkFn)

	return entry
}
