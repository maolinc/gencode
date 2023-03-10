package filex

import (
	"embed"
	"golang.org/x/mod/modfile"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func AppendToFile(filename string, byte []byte) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		return err
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, io.SeekEnd)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt(byte, n)
	}
	defer f.Close()
	return err
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

func FindFileToBack(path, fileName string) (context []byte, destPath string, err error) {
	if filepath.IsAbs(path) {
		return findFile(path, fileName)
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, "", err
	}
	return findFile(absPath, fileName)
}

func findFile(path string, filename string) (context []byte, destPath string, err error) {
	fullPath := filepath.Join(path, filename)
	exist, err := PathExists(fullPath)
	if err != nil {
		return nil, "", err
	}
	if exist {
		if file, err := os.ReadFile(fullPath); err == nil {
			return file, path, nil
		}
		return nil, "", err
	}
	parentDir := filepath.Dir(path)
	if path == parentDir {
		return nil, parentDir, nil
	}
	return findFile(parentDir, filename)
}

func GetAbs(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return abs
}

func GetModule(path string) (module, fixPath string) {
	var (
		modName  = "go.mod"
		context  []byte
		destPath string
		err      error
	)
	if !filepath.IsAbs(path) {
		if path, err = filepath.Abs(path); err != nil {
			return "", ""
		}
	}
	context, destPath, err = findFile(path, modName)
	if err != nil {
		return "", ""
	}
	modFile := modfile.ModulePath(context)
	path = strings.TrimPrefix(path, destPath)
	path = strings.ReplaceAll(path, "\\", "/")
	return modFile, path
}
