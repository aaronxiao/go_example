package common

import (
	"os"
	"path/filepath"
)

//获得当前目录路径
func GetCurPwd() string  {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	dir := filepath.Dir(path)
	//fmt.Printf("GetCurPwd path: %s dir: %s \n", path, dir)
	return dir
}
