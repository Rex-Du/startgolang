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
	"strings"
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
		//meta.UpdateFileMetas(fileMeta)
		meta.UpdateFileMetaDB(fileMeta)
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

func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	optType := r.Form.Get("opt") // 操作类型，0是重命名，1是删除
	filehash := r.Form.Get("filehash")
	if optType == "0" {
		newName := r.Form.Get("newname")
		if newName == "" {
			io.WriteString(w, "重命名文件需要传入新的文件名newname参数")
			return
		}
		fmeta := meta.GetFileMeta(filehash)
		newPath := strings.Replace(fmeta.Location, fmeta.FileName, newName, -1)
		err := os.Rename(fmeta.Location, newPath)
		if err != nil {
			panic(err.Error())
		}
		fmeta.FileName = newName
		fmeta.Location = newPath
		//meta.UpdateFileMetas(fmeta)
		meta.UpdateFileMetaDB(fmeta)
		io.WriteString(w, "重命名成功")
	}
}

func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form.Get("filehash")
	fmeta := meta.GetFileMeta(fileHash)

	os.Remove(fmeta.Location)

	meta.DeleteFileMeta(fileHash)
	w.WriteHeader(http.StatusOK)
}
