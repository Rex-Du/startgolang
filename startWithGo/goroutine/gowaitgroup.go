package main

import (
	"fmt"
	"sync"
)

/*
waitGroup让主函数等待子goroutine执行完毕
 */

var wg sync.WaitGroup
func main()  {
	wg.Add(2)
	go func1()
	go func2()

	fmt.Println("主函数进入等待状态。。。")
	wg.Wait()
	fmt.Println("主函数结束。。。")

}

func func1()  {
	for i:=0;i<10;i++{
		fmt.Printf("子函数1打印%d\n",i)
	}
	wg.Done()

}

func func2()  {
	for i:=0;i<10;i++{
		fmt.Printf("\t子函数2打印%d\n",i)
	}
	wg.Done()
}