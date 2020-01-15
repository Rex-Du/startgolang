package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":9090")
	if err != nil {
		panic(err.Error())
	}
	var res float32
	res = 3
	var resp *float32
	err = client.Call("MathUtil.CalculateCircleArea", res, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)
}
