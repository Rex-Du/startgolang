package main

import "fmt"

func main()  {
	s := "YES我爱go语言"

	for i, ch := range []rune(s){
		fmt.Printf("(%d %c) ", i, ch)
	}

}
