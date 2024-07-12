package game

import (
	"os"
	"path/filepath"
	"strings"
)

func listDir(dir string) (filePaths []string, baseNames []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 遇到错误时返回
		}
		if !info.IsDir() { // 确保是文件
			filePaths = append(filePaths, path)                                     // 添加完整路径到切片
			baseName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) // 去除文件后缀
			baseNames = append(baseNames, baseName)                                 // 添加去除后缀的文件名到切片
		}
		return nil
	})

	if err != nil {
		return nil, nil, err // 遇到错误时返回
	}

	return filePaths, baseNames, nil
}
