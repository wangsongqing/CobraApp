package controllers

import (
	"CobraApp/pkg/logger"
	"CobraApp/pkg/redis"
	"fmt"
	"github.com/gookit/color"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type ChannelTest struct {
	RedisKey string
}

func (c *ChannelTest) Push() {

	for i := 0; i < 10; i++ {
		res := redis.Redis.Rpush(c.RedisKey, i)
		if res == false {
			color.Redln("写入失败")
		}
	}

	color.Redln("写入成功")
}

// Pop 单个协程消费
func (c *ChannelTest) Pop() {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		for {
			data := redis.Redis.Lpop(c.RedisKey)
			if len(data) == 0 {
				time.Sleep(time.Second)
				continue
			}
			time.Sleep(time.Second)
			fmt.Printf("data:%v \n", data)
		}

		// 如果不是常驻job需要done协程减1
		// w.Done()
	}()
	w.Wait()
}

// MultiProcessingPop 多个协程消费队列
func (c *ChannelTest) MultiProcessingPop() {
	var w sync.WaitGroup

	// 获取CPU核数，按CPU数量来开协程数量是最合适的
	cpuNum := runtime.NumCPU()
	w.Add(cpuNum)
	for i := 0; i < cpuNum; i++ {
		go func() {
			for {
				data := redis.Redis.Lpop(c.RedisKey)
				if len(data) == 0 {
					time.Sleep(time.Second)
					continue
				}

				time.Sleep(time.Second)
				fmt.Printf("data:%v \n", data)
				res, _ := strconv.Atoi(data)
				if res%2 == 0 {
					logger.Info("打印取模等于0的数据:" + data)
				}
			}

			// 如果不是常驻job需要Done协程减1
			// w.Done()
		}()
	}

	w.Wait()
}
