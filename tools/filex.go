package tools

import (
	"embed"
	"io"
	"os"
)

func GetHomeDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir
}

func CopyDirEm(dir embed.FS, toDir string, file string) (err error) {
	readDir, err := dir.ReadDir(file)
	if err != nil {
		return err
	}

	err = os.MkdirAll(toDir, 777)
	if err != nil {
		return err
	}

	for _, entry := range readDir {
		fileTo, err := os.OpenFile(toDir+"/"+entry.Name(), os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		readFile, err := dir.ReadFile(file + "/" + entry.Name())
		_, err = io.WriteString(fileTo, string(readFile))
		if err != nil {
			return err
		}
		fileTo.Close()
	}
	return nil
}

func CopyDir(fromDir, toDir string) (err error) {
	fileFrom, err := os.Open(fromDir)
	if err != nil {
		return err
	}
	dirs, err := fileFrom.ReadDir(-1)
	if err != nil {
		return err
	}

	err = os.MkdirAll(toDir, 777)
	if err != nil {
		return err
	}

	for _, file := range dirs {
		openFile, err := os.OpenFile(fromDir+"/"+file.Name(), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		fileTo, err := os.OpenFile(toDir+"/"+file.Name(), os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = io.Copy(fileTo, openFile)
		if err != nil {
			return err
		}
		fileTo.Close()
		openFile.Close()
	}

	defer fileFrom.Close()

	return err
}

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
