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

func (t *Test) ChannelPushList() {
	redis.Redis.Rpush(key, "111")
}

// ChannelPopList 协程消费队列
func (t *Test) ChannelPopList() {
	for {
		data := redis.Redis.Lpop(key)
		if len(data) == 0 {
			time.Sleep(time.Second)
			continue
		}
		for i := 1; i < 3; i++ {
			go func(is int) {
				fmt.Println(data)
				time.Sleep(time.Second * 10)
				fmt.Println(data + "_" + strconv.Itoa(is))
			}(i)
		}

	}

}
