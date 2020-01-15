package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


//FileUploadHandler xxx
func FileUploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET"{
		data, _ := ioutil.ReadFile("static/index.html")
		_, err :=io.WriteString(w, string(data))
		if err != nil{
			log.Print(err.Error())
		}
	}else if r.Method == "POST"{
		file, head, err := r.FormFile("file")
		if err != nil{
			panic(err.Error())
		}
		new_file, err := os.Create("/tmp/"+head.Filename)
		if err != nil{
			fmt.Printf("create file failed: %s", err.Error())
		}
		defer new_file.Close()

		_, err = io.Copy(new_file, file)
		if err != nil{
			panic(err.Error())
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "upload success!")
}