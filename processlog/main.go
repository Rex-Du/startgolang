package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type processLog struct {
	reader Reader
	writer Writer
	rc chan []byte
	wc chan string
}

func (pl *processLog) ReadData()  {
	pl.reader.Read(pl.rc)
}

func (pl *processLog) HandleData()  {
	for data := range pl.rc{
		pl.wc <- strings.ToUpper(string(data))
	}

}

func (pl *processLog) WriteData()  {
	for data := range pl.wc{
		pl.writer.Write(data)
	}
}

type Reader interface {
	Read(ch chan []byte)
}

type Writer interface {
	Write(string)
}

type FileReader struct {
	filePath string
}

type DBWriter struct {

}

func (fr FileReader) Read(ch chan []byte)  {
	// 从文件读取数据
	//data := "asdfghjkl"
	f, err := os.Open(fr.filePath)
	if err != nil{
		panic(fmt.Sprintf("open file error: %s", err.Error()))
	}
	defer f.Close()
	//f.Seek(0, 2)

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		if err != nil{
			if err == io.EOF{
				time.Sleep(500*time.Millisecond)
				continue
			}else {
				panic(err.Error())
			}
		}
		ch <- line
	}
	//rd := bufio.NewScanner(f)
	//for rd.Scan(){
	//	ch <- []byte(rd.Text())
	//}
}

func (dbw DBWriter) Write(data string)  {
	fmt.Print(data)
}

func main() {
	pl := &processLog{reader:FileReader{filePath:"./access.log"}, writer:DBWriter{}, rc:make(chan []byte), wc:make(chan string)}
	go pl.ReadData()
	go pl.HandleData()
	go pl.WriteData()
	time.Sleep(100*time.Second)
}
