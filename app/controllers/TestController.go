package controllers

import (
	"CobraApp/pkg/redis"
	"fmt"
	"strconv"
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
