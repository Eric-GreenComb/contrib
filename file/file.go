package file

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	bgstrings "github.com/banerwai/gommon/strings"
)

// SplitFilename 分离文件名与扩展名(包含.)
func SplitFilename(filename string) (baseName, ext string) {
	baseName = filename
	// 找到最后一个'.'
	ext = bgstrings.SubstringByte(filename, strings.LastIndex(filename, "."))
	baseName = strings.TrimRight(filename, ext)
	ext = strings.ToLower(ext)
	return
}

// TransferExt 转换文件的格式
// toExt包含.
func TransferExt(path string, toExt string) string {
	dir := filepath.Dir(path) + "/" // 文件路径
	name := filepath.Base(path)     // 文件名 a.jpg
	// 获取文件名与路径
	baseName, _ := SplitFilename(name)
	return dir + baseName + toExt
}

// GetFilename get file name
func GetFilename(path string) string {
	return filepath.Base(path)
}

// GetFilesize get file size
// length in bytes
func GetFilesize(path string) int64 {
	fileinfo, err := os.Stat(path)
	if err == nil {
		return fileinfo.Size()
	}
	return 0
}

// CopyFile cp file
func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// DeleteFile delete file
func DeleteFile(path string) bool {
	err := os.Remove(path)
	if err != nil {
		return false
	}
	return true
}

// GetFileStrContent 获得文件str内容
func GetFileStrContent(path string) string {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(fileBytes)
}

// IsFileExist check file is exist
func IsFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// PutFileStrContent 写入string内容
func PutFileStrContent(path, content string) bool {
	var f *os.File
	var err1 error
	defer (func() {
		if f != nil {
			f.Close()
		}
	})()
	f, err1 = os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
	//	Log(err1)
	//	var n int
	_, err1 = io.WriteString(f, content) //写入文件(字符串)
	//	Log(content)
	//	Log(err1)
	//	Log(n)
	//	Log(path)

	if err1 != nil {
		return false
	}
	return true
}
