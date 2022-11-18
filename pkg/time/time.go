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

// TimesToStamp TimeToStamp 时间 to 时间戳
func TimesToStamp(times string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")                     //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", times, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	return tt.Unix()
}

// StampToTime 时间戳 to 时间
func StampToTime(strTime int64) string {
	tm := time.Unix(strTime, 0)
	return tm.Format("2006-01-02 15:04:05") //2018-07-11 15:10:19
}
