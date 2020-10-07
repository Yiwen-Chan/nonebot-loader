package global

import (
	"io"
	"net/http"
)

func DownFile(url, file string) (err error) {
	INFO("正在下载数据", url)
	INFO("目标文件路径", file)
	res, err := http.Get(url)
	if err != nil {
		ERROR("下载文件失败", url)
		return err
	}
	f, err := CreateFile(file)
	if err != nil {
		ERROR("创建文件失败")
		return err
	}
	_, err = io.Copy(f, res.Body)
	if err != nil {
		ERROR("写入文件失败")
		return err
	}
	INFO("下载数据完毕", url)
	return nil
}
