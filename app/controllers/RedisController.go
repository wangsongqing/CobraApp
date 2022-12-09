package controllers

import (
	"CobraApp/pkg/redis"
	"fmt"
	"github.com/gookit/color"
	"sync"
	"time"
)

type RedisTest struct {
}

func (r *RedisTest) Lock() {

	lockName := "my_test_lock"

	// 2秒钟没有拿到锁，说明锁已经被占用了，锁的过期时间为60
	acquired, err := redis.Redis.AcquireLock(lockName, 2*time.Second, 60*time.Second)
	if err != nil {
		return
	}

	if acquired == false {
		fmt.Println("已经有程序在处理了")
		return
	}

	fmt.Println("Successfully acquired lock")
	// 在这里执行业务逻辑

	time.Sleep(time.Second * 10)
	color.Redln("处理一些数据成功")

	// 业务逻辑执行完毕后，释放锁
	released, err := redis.Redis.ReleaseLock(lockName)
	if err != nil {
		fmt.Println(err)
		return
	}

	if released {
		fmt.Println("Successfully released lock")
	} else {
		fmt.Println("Failed to release lock")
	}
}

// 全局变量
var counter int

func (r *RedisTest) LockApplication() {
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}
