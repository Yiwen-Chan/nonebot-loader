package main

import (
	"nonebot-loader/global"
)

func main() {
	python()
}

func python() {
	path := global.PathExecute()

	url := "https://www.python.org/ftp/python/3.8.6/python-3.8.6-embed-amd64.zip"
	file := path + "temp/python.zip"

	env := path + "python/"
	py := env + "python.exe"

	pth := env + "python38._pth"
	pipUrl := "https://uploader.shimo.im/f/eVYoo89X7YQ2wo6x.py"
	pipFile := env + "get-pip.py"

	mirror := "http://mirrors.aliyun.com/pypi/simple/"

	pipExe := env + "Scripts/pip3.exe"

	sitecustomize := env + "Lib/site-packages/sitecustomize.py"

	nonebot := env + "Lib/site-packages/nonebot/"
	pipTxt := "nonebot"

	//下载 python embedded 环境
	exists, err := global.PathExists(py)
	if err != nil {
		return
	}
	if !exists {

		//下载 python embedded 并解压
		global.DownFile(url, file)
		global.UnpackZIP(file, env)
	}

	//下载 pip
	exists, err = global.PathExists(pipFile)
	if !exists {
		//修改 python embedded 设置 并下载pip
		global.WriteFile(pth, "python38.zip\n.\n..\nimport site")
		global.DownFile(pipUrl, pipFile)

	}

	//安装 pip
	exists, err = global.PathExists(pipExe)
	if !exists {
		global.Cmd(py, pipFile)
	}

	//修复 python 在 exec 运行时编码错误
	exists, err = global.PathExists(sitecustomize)
	if !exists {
		global.WriteFile(sitecustomize, "import io\nimport sys\nsys.stdout = io.TextIOWrapper(sys.stdout.buffer,encoding='gb18030')")
	}

	//安装 nonebot 环境
	exists, err = global.PathExists(nonebot)
	if !exists {
		global.Cmd(pipExe, "install", pipTxt, "-i", mirror)
	}
}
