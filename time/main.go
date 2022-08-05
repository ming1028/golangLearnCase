package main

import (
	"fmt"
	"time"
)

// GMT格林尼治时间，世界标准时间 UTC协调世界时，0时区时间 中国东八区 UTC+08
var timeStr = "2006-01-02 15:04:05"

func main() {
	now := time.Now()
	nowStr := now.Format("2006年01月02日 15:04:05")
	nowStr = now.Format(timeStr)
	fmt.Println(nowStr)
	fmt.Println(now.Format(timeStr))
	fmt.Println(now.Format(time.Layout), now.Format(time.UnixDate))
	timestamp := now.Unix()
	fmt.Println("时间戳:", timestamp)
	fmt.Println("时间戳转时间:", time.Unix(timestamp, 0))
	nowTime, _ := time.Parse(timeStr, nowStr)
	fmt.Println("字符串转时间类型", nowTime)

	fmt.Println("时间加减")

	dura, err := time.ParseDuration("10h")
	fmt.Println(dura, err)
	fmt.Println("10小时之后：", now.Add(dura))

	dura, err = time.ParseDuration("-10h")
	fmt.Println(dura, err)
	fmt.Println("10小时之前：", now.Add(dura))

	start := time.Now()
	time.Sleep(5)
	end := time.Now()
	fmt.Println("耗时：", end.Sub(start))

	fmt.Printf("秒：Time type:%T, value:%v\n", now.Unix(), now.Unix())
	fmt.Printf("纳秒：type:%T, value:%v\n", now.UnixNano(), now.UnixNano())
	fmt.Println(now.Unix(), now.UnixMicro(), now.UnixNano(), now.UnixMilli())
	fmt.Println(time.Hour*5 + time.Second*5 + time.Minute*5)

	targetTime := now.Add(time.Hour * 5)
	fmt.Println(targetTime.Sub(now))
	fmt.Println("时间相同：", now.Equal(targetTime))
	fmt.Println("时间是否在之前：", now.Before(targetTime))
	fmt.Println("时间是否在之后：", now.After(targetTime))

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	fmt.Println(secondsEastOfUTC)
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Println(time.Date(2006, 01, 02, 15, 04, 05, 0, beijing))
	fmt.Println(time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local))
	fmt.Println(time.Now().In(beijing))

	date, err := time.ParseInLocation(timeStr, "2020-01-28 11:05:06", time.Local)
	fmt.Println(date, err)

	fmt.Println("当前时间在零时区时间：", now.UTC().Year(), now.UTC().Month(), now.UTC().Day(),
		now.UTC().Hour(), now.UTC().Minute(), now.UTC().Second(),
	)

	hour, min, sec := now.UTC().Clock()
	fmt.Println("当前时间在零时区,时分秒：", hour, min, sec)

	fmt.Println("定时器")
	ticker := time.Tick(time.Second)
	for v := range ticker {
		fmt.Println(v.Second())
	}

}
