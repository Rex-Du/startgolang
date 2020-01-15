package main

import "fmt"

//func main(){
//	var a []int = []int{4, 5,6}
//	fmt.Println(a)
//	s := append(a, 10)
//	fmt.Println(s)
//}
func main() {
	var a []int
	for i := 0; i < 100; i++ {
		printSlice(a)
		a = append(a, i)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
