package main

import "fmt"

/**
go中的闭包
*/

func main()  {
	f1 := increment()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

}

func increment() func() int {
	i := 0

	add := func() int {
		i++
		return i
	}
	//func add() int{
	//	i++
	//	return i
	//}
	return add

}
