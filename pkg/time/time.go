package time

import (
	"fmt"
	"time"
)

// NowTimeDate 获取时间(格式话获取时间)
func NowTimeDate() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format("2006-01-02 15:04:05"))
}

// GetTimeNowDate 只获取年月日
func GetTimeNowDate() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format("2006-01-02"))
}

// GetNowTime 获取时间(时间戳)
func GetNowTime() int64 {
	return time.Now().Unix()
}
