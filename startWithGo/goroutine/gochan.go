package main

import "fmt"

func main()  {
	var ch chan int		//声明一个chan
	ch = make(chan int)		//给chan赋值
	go func() {
		for i:=0;i<10;i++{
			fmt.Println(i)
		}
		ch <- 1
	}()

	data := <-ch
	fmt.Println(data)


}