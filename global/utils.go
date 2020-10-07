package global

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func PathExecute() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)

	return dir + "/"
}

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

func GetDir(file string) string {
	return subString(file, 0, strings.LastIndex(file, "/"))
}

func CreateFile(file string) (io.Writer, error) {
	err := os.MkdirAll(getDir(file), 0755)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func WriteFile(file, content string) {
	INFO("正在写入文件", file)
	exists, _ := PathExists(file)
	if exists {
		os.Remove(file)
	}
	f, err := os.OpenFile(file, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		ERROR("打开文件失败")
		return
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		ERROR("写入文件失败")
		return
	}
	INFO("写入文件完毕", file)
}
