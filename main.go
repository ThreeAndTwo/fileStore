package main

import (
	"fileStore/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	err := http.ListenAndServe(":8080", nil)

	if nil != err {
		fmt.Printf("Failed to start server, err: %s", err.Error())
	}

}
