package mylib

import (
	"os/exec"
	"syscall"
)

//执行cmd不弹窗，中文无乱码
func Openfile(path string) {
	cmd := exec.Command(`cmd`, `/c`, `start`, path)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()

}
