package mylib

import (
	"errors"
	"github.com/CodyGuo/win"
	"log"
	"os/exec"
	"syscall"
)

var (
	winExecError = map[uint32]string{
		0:  "The system is out of memory or resources.",
		2:  "The .exe file is invalid.",
		3:  "The specified file was not found.",
		11: "The specified path was not found.",
	}
)

//执行cmd不弹窗，中文无乱码
func Openfile(path string) {
	cmd := exec.Command(`cmd`, `/c`, `start`, path)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

}

func Cmdstart(path1 string) {
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
