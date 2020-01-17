package handlers

import (
	"clouddisk/meta"
	"clouddisk/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//FileUploadHandler xxx
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, _ := ioutil.ReadFile("static/index.html")
		_, err := io.WriteString(w, string(data))
		if err != nil {
			log.Print(err.Error())
		}
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")

		fileMeta := meta.FileMeta{
			FileSha1: "",
			FileName: head.Filename,
			FileSize: 0,
			Location: "tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		if err != nil {
			panic(err.Error())
		}
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("create file failed: %s", err.Error())
		}

		defer newFile.Close()

		if err != nil {
			panic(err.Error())
		}

		fileMeta.FileSize, err = io.Copy(newFile, file)
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMetas(fileMeta)
		//http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
		resp, _ := json.MarshalIndent(fileMeta, "", "\t")
		io.WriteString(w, string(resp))
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload success!")
}

func FileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//name := r.Form.Get("filename")
	//file, _ := os.Open(name)
	//filehash := util.FileSha1(file)
	filehash := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(filehash)
	resp, _ := json.MarshalIndent(fileMeta, "", "\t")
	io.WriteString(w, string(resp))
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fmeta := meta.GetFileMeta(fsha1)
	f, err := os.Open(fmeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=\""+fmeta.FileName+"\"")
	w.Write(data)

}
