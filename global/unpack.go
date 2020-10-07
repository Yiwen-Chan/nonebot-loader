package global

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func UnpackZIP(zipFile, dest string) error {
	INFO("开始解压", zipFile)
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		ERROR("解压出错", zipFile)
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		if IsFile(filename) {
			w, err := os.Create(filename)
			if err != nil {
				ERROR("解压过程出错", filename)
				return err
			}
			defer w.Close()
			_, err = io.Copy(w, rc)
			if err != nil {
				return err
			}
			w.Close()
		} else {
			os.Mkdir(filename, os.ModePerm)
		}
		rc.Close()
	}
	INFO("解压完毕", zipFile)
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func IsFile(f string) bool {
	last := f[len(f)-1:]
	if last == "/" {
		return false
	}
	return true
}
