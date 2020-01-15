package main

import "fmt"

func adder() func(int) int  {
	sum :=0
	return func(value int) int{
		sum += value
		return sum
	}
}

func main()  {
	a := adder()
	for i:=0;i<10;i++{

		fmt.Printf("1+2+...+%d=%d\n", i, a(i))
	}
}


