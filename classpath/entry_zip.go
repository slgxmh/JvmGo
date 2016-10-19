package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)
/**
表示ZIP或JAR文件的类路径
 */
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

//返回读取结果
func (self *ZipEntry)readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	//若打开文件失败直接返回
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	//遍历整个文件夹，寻找文件
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

//返回路径
func (self *ZipEntry)String() string {
	return self.absPath
}