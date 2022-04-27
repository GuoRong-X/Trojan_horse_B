package mylib

import (
	"fmt"
	"os"
)

//正常退出
func Stop() {
	fmt.Println("结束")
	os.Exit(0)
}
