package main

import (
	"fmt"
	"piplinesort/pipline"
)

func main() {
	p1 := pipline.ArrayIn(3, 8, 10, 2, 4)
	middle1 := pipline.SortIn(p1)
	p2 := pipline.ArrayIn(1, 5, 7, 6, 9)
	middle2 := pipline.SortIn(p2)

	out := pipline.Merge(middle1,middle2)
	for v := range out {
		fmt.Println(v)
	}
}
