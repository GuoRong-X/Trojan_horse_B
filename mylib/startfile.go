package mylib

import (
	"errors"
	"log"
)
import (
	"github.com/CodyGuo/win"
)

//暂停使用，中文乱码
//调用cmd打开文件，并实现不弹doc窗口
var (
	winExecError = map[uint32]string{
		0:  "The system is out of memory or resources.",
		2:  "The .exe file is invalid.",
		3:  "The specified file was not found.",
		11: "The specified path was not found.",
	}
)

func Cmdpath(path1 string) {
	var cmdpath = "cmd /c start " + path1
	err := execRun(cmdpath)
	if err != nil {
		log.Fatal(err)
	}
}

func execRun(cmd string) error {

	lpCmdLine := win.StringToBytePtr(cmd)
	ret := win.WinExec(lpCmdLine, win.SW_HIDE)
	if ret <= 31 {
		return errors.New(winExecError[ret])
	}
	return nil
}
