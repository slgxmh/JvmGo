package classpath

import (
	"path/filepath"
	"io/ioutil"
)
/**
表达目录形式的类路径
 */
type DirEntry struct {
	//绝对地址
	absDir string
}

//创建一个文件实体类
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

//拼接为文件完整路径，返回读取结果
func (self *DirEntry) readClass(classname string) ([]byte, Entry, error) {
	filename := filepath.Join(self.absDir, classname)
	data, err := ioutil.ReadFile(filename)
	return data, self, err
}

//返回绝对路径
func (self *DirEntry)String() string {
	return self.absDir
}