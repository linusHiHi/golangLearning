package main

import (
	"fmt"
	"time"
)

func main() {
	//time zone
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		return
	}
	beginTime := time.Now()
	fmt.Println(beginTime.In(loc))

	//using clock() to get time
	hours, mins, secs := beginTime.Clock()
	fmt.Printf("程序开始时间：%d:%d:%d\n", hours, mins, secs)
	//hour()...等效方法
	fmt.Printf("程序开始时间：%d:%d:%d\n", beginTime.Hour(), beginTime.Minute(), beginTime.Second())
	//获取日期
	year, month, day := beginTime.Date()
	fmt.Printf("date: %d,%d,%d\n", year, month, day)
	fmt.Printf("%d,%d,%d\n", beginTime.Year(), beginTime.Month(), beginTime.Day())

	//time addition
	twoHourLater := beginTime.Add(time.Hour * 2)
	fmt.Println("two hour later: ", twoHourLater)
	someDateLater := beginTime.AddDate(1, 3, 3)
	fmt.Println("later", someDateLater)

	//duration
	//set duration
	midTime := time.Now()
	duration := midTime.Sub(beginTime)
	println(duration)
	durationSince := time.Since(beginTime)
	println(durationSince)

	//截断精度
	truncatedTime := beginTime.Truncate(time.Hour)
	fmt.Printf("%v\n", truncatedTime)

	preFormatString := time.Now()
	formatString := preFormatString.Format("2006-02-01 15:04:05")
	fmt.Println("yyyy-dd-mm hh:mm:ss", formatString)

	// 解析字符串为时间
	timeStr := "2023-11-15 12:30:45"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("解析时间出错：", err)
	} else {
		fmt.Println("解析后的时间：", parsedTime)
	}

	// 解析时间段字符串
	durationStr := "1h30m45s"
	parsedDuration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Println("解析持续时间出错：", err)
	} else {
		fmt.Printf("解析后的持续时间：%v，总小时数：%f\n", parsedDuration, parsedDuration.Hours())
	}

	//timer &&ticker
	timerPrinter := time.NewTimer(durationSince)
	defer timerPrinter.Stop()
	for i := 1; i >= 5; i++ {
		<-timerPrinter.C // 等待定时器通道触发

		// 在这里执行定时任务
		fmt.Println("timer at", time.Now())
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C // 等待定时器通道触发

		// 在这里执行定时任务
		fmt.Println("Tick at", time.Now())
	}

}
