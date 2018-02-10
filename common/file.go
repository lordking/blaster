package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//ReadFileData 读取文件所有内容
func ReadFileData(path string) ([]byte, error) {
	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadFile(path); err != nil {
		return nil, NewError(ErrCodeInternal, err.Error())
	}

	return data, nil
}

//AppendIntoFile 添加数据到文件最后位置
func AppendIntoFile(data []byte, filePath string) error {

	var err error

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	info, err := f.Stat()
	if err != nil {
		return err
	}

	_, err = f.WriteAt(data, info.Size())
	if err != nil {
		return err
	}

	return nil
}

//replaceHelper 文件查找替换器定义
type replaceHelper struct {
	Suffix  string
	Regexp  *regexp.Regexp
	NewText []byte
}

//ReplaceInDirectory 批量替换文件
func ReplaceInDirectory(directory, suffix, regx, newText string) error {
	r := &replaceHelper{suffix, regexp.MustCompile(regx), []byte(newText)}
	return filepath.Walk(directory, r.walkCallback)
}

func (t *replaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if f == nil {
		return nil
	}

	if f.IsDir() {
		return nil
	}

	//后缀名不符合的文件，不替换
	if !strings.HasSuffix(f.Name(), t.Suffix) {
		return nil
	}

	//读取文件
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	//替换
	newBuf := t.Regexp.ReplaceAll(buf, t.NewText)

	//重新写入
	ioutil.WriteFile(path, newBuf, 0)

	return err
}
