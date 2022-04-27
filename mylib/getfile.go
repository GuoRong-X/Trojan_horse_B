package mylib

import (
	"io"
	"net/http"
	"os"
)

//从云端下载文件，并保存到本地
func DownloadFile(url string, path string) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create output file
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// copy stream
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
