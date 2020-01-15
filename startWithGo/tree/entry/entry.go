package main

import (
	"fmt"
	"startWithGo/tree"
)

func main()  {
	root := tree.Node{Value:3}
	root.Left = &tree.Node{Value:5}
	root.Print()
	root.Left.Print()

	root.SetValue(10)
	root.Left.SetValue(20)
	root.Print()
	root.Left.Print()
	fmt.Println(root)

}
