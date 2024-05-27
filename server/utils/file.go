package utils

import (
	"errors"
	"os"
)

/*
包含目录和文件相关的所有工具函数
*/

// PathExists 判断目录是否存在
// 1.若不存在,则返回false,nil
// 2.若存在且为目录,则返回true,nil
//
//tips:error不为空的情况, 根据实际使用情况判断是否输出error内容或panic
func PathExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) { // 判断是否是文件不存在导致的错误
			return false, nil
		} else {
			return false, err
		}
	} else {
		if file.IsDir() { // 判断文件是否是目录
			return true, nil
		} else {
			return false, errors.New("the file has same name is exist")
		}
	}
}
