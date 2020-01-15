package main

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type MathUtil struct {
}

func (mu *MathUtil) CalculateCircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req //返回圆形的面积
	for i:=0;i<10;i++{
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
	return nil
}

func main()  {
	mu := new(MathUtil)
	err :=rpc.Register(mu)
	if err !=nil{
		panic(err.Error())
	}

	listen, err := net.Listen("tcp", ":9090")
	if err != nil{
		panic(err.Error())
	}

	rpc.HandleHTTP()

	http.Serve(listen, nil)
}