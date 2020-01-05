package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传 html 页面
		fileData, fileErr := ioutil.ReadFile("./static/view/index.html")

		if nil != fileErr {
			// 文件读取错误
			io.WriteString(w, "internel Server error")
			return
		}
		io.WriteString(w, string(fileData))

	} else if r.Method == "POST" {
		// 接收文件流及存储到本地目录
		r.FormFile("")

	}
}
