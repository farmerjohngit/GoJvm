package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry struct {
	entries []Entry
}

func newCompositeEntry(path string) *CompositeEntry {
	compoundEntry := &CompositeEntry{}
	for _, str := range strings.Split(path, pathListSeparator) {
		compoundEntry.addEntry(newEntry(str))
	}
	return compoundEntry
}

func (self *CompositeEntry) readClass(className string) ([] byte, error) {
	for _, entry := range self.entries {
		data, error := entry.readClass(className)
		println("find class in ",entry.toString())
		if error == nil {
			return data, error
		}
	}
	return nil, errors.New("class not found: " + className)
}

func (self *CompositeEntry) addEntry(entry Entry) {
	self.entries = append(self.entries, entry)
}

func (self *CompositeEntry) toString() string {
	stringArr := make([]string, len(self.entries))
	for index, entry := range self.entries {
		stringArr[index] = entry.toString()
	}
	return strings.Join(stringArr, pathListSeparator)
}
