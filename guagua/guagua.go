package main

import "github.com/kataras/iris"

import "github.com/kataras/iris/mvc"

import "time"

import "math/rand"

import "strconv"

import "sort"

type lotteryController struct {
	Ctx iris.Context
}

// NewApp xxx
func NewApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := NewApp()
	app.Run(iris.Addr(":8000"))
}

// 刮刮乐，1 一等奖，2-3 二等奖，4-6 三等奖，7-10无奖
func (c *lotteryController) Get() string  {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Intn(10)
	var prize string
	switch {
	case code == 1:
		prize = "一等奖"
	case code == 2 || code == 3:
		prize = "二等奖"
	case code >=4 && code <= 6:
		prize = "三等奖"
	default:
		return "你没中奖"
	}
	return "你中了" + prize
}

// 双色球 直接显示当期中奖的号码
func (c *lotteryController) GetDouble() string{
	
	// r := rand.New(rand.NewSource()).Read()
	r := genRandomNumber(33, 6)
	b := genRandomNumber(16, 1)
	res :="本期中奖号码是："
	for _, v := range r{
		res += strconv.Itoa(v) + " "
	}
	res += strconv.Itoa(b[0])
	return res
}

func genRandomNumber(n, count int) []int{
	var res = make([]int, 0)
	if n < count{
		return nil
	}
	for {
		if len(res) < count{
			seed := time.Now().UnixNano()
			code := rand.New(rand.NewSource(seed)).Intn(n) + 1

			exist := false
			for _, v := range res{
				if v == code{
					exist = true
					break
				}
			}
			if !exist{
				res = append(res, code)
			}
		}else{
			break
		}
	}
	sort.Ints(res)
	return res

}
