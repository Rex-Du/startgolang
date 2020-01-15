package main

import (
	"clouddisk/handlers"
	"fmt"
	"net/http"
)


func main()  {
	http.HandleFunc("/file/upload", handlers.FileUploadHandler)
	http.HandleFunc("/file/upload/suc", handlers.UploadSucHandler)
	err := http.ListenAndServe(":8989", nil)
	if err != nil{
		fmt.Println("start server error")
	}
}