package main

import (
	"fmt"
	"time"
)

func main()  {
	go printNum()
	for i:=0;i<10000;i++{
		fmt.Printf("\t主goroutine打印字母：A %d\n",i)
	}
	time.Sleep(1*time.Second)
	fmt.Println("main。。。。over。。。。")
}

func printNum()  {
	for i:=0;i<10000;i++{
		fmt.Printf("子goroutine打印数字：%d\n", i)
	}

}
