package util

import "strings"

type FileUtil struct {
}

func IsArchive(path string) bool {
	return isSomeSuffixFile("path", ".jar", ".zip")
}

func IsClassFile(path string) bool {
	return isSomeSuffixFile(path, "class")
}

func isSomeSuffixFile(path string, suffixes ... string) bool {
	lowerPath := strings.ToLower(path)
	for _, suffix := range suffixes {
		if strings.HasSuffix(lowerPath, suffix) {
			return true
		}
	}
	return false
}
