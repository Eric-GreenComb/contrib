package file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Substr golang substr
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// GetParentDirectory get parent directory
func GetParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

// GetCurrentDirectory get current directory
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// ClearDir 清空dir下所有的文件和文件夹
// RemoveAll会清空本文件夹, 所以还要创建之
func ClearDir(dir string) bool {
	err := os.RemoveAll(dir)
	if err != nil {
		return false
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return false
	}
	return true
}

// MkdirAll 创建文件夹
func MkdirAll(dir string) bool {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return false
	}
	return true
}

// ListDir list dir's all file, return filenames
func ListDir(dir string) []string {
	f, err := os.Open(dir)
	if err != nil {
		return nil
	}
	names, _ := f.Readdirnames(0)
	return names
}

// CopyDir cp dir
func CopyDir(source string, dest string) (err error) {
	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				//				fmt.Println(err)
			}
		} else {
			// perform copy
			_, err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				//				fmt.Println(err)
			}
		}
	}
	return
}

// IsDirExists check if dir is exist
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return fi.IsDir()
}
