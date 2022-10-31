package controllers

import (
	"CobraApp/pkg/redis"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Test struct {
}

const key = "test_list__cobraapp"

func (st *Test) ChannelPushList() {
	redis.Redis.Rpush(key, "keys:")
}

// ChannelPopList 协程消费队列
func (st *Test) ChannelPopList() {
	for {
		data := redis.Redis.Lpop(key)
		if len(data) == 0 {
			time.Sleep(time.Second)
			continue
		}

		go func() {
			//fmt.Println(data)
			time.Sleep(time.Second * 10)
			fmt.Println(data + "_" + strconv.Itoa(0))
		}()

	}
}

func (st *Test) ChangeName(name *string) {
	*name = *name + " hello"
}

func (st *Test) SliceTest() {
	slices := []any{1, 2, 3, "ww"}
	for _, value := range slices {
		fmt.Printf("type:%T, value:%v \n", value, value)
	}

}

var w sync.WaitGroup

// Chanelle 开启多了协程处理程序
func (st *Test) Chanelle() {

	cpu := runtime.NumCPU() // 获取机器的CPU数量

	// 开启10个协程处理队列
	for i := 0; i < cpu; i++ {
		w.Add(1) // 协程计数器加1
		//go func(n int) {
		//	num := helpers.MakeRandInt()
		//	fmt.Printf("协程%v, 随机数:%v \n", n, num)
		//	time.Sleep(time.Second)
		//	w.Done()
		//}(i)

		go PopData()
	}

	w.Wait() // 等待所有的协程执行完毕
}

func PopData() {
	for {
		data := redis.Redis.Lpop(key)
		if len(data) == 0 {
			time.Sleep(time.Second)
			continue
		}

		fmt.Printf("pop_data:%v \n", data)
		time.Sleep(time.Second)
	}

	// 如果不是常驻job需要done协程减1
	// w.Done() // 协程计数器减1
}
