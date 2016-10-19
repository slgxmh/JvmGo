package classpath

import (
	"strings"
)
/**
复合地址
 */
type CompositeEntry[] Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

//遍历读取每一个类
func (self CompositeEntry) readClass(className string) ([] byte, Entry, error) {
	strs := make([]string, len(self))
	for i, entey := range self {
		strs[i] = entey.String()
	}
	return strings.Join(strs, pathListSeparator)
}

func (self CompositeEntry)String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}