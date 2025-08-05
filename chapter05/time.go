package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前本地时间
	now := time.Now()
	fmt.Printf("当前时间: %v\n", now)
	// 获取UTC时间
	utc := now.UTC()
	fmt.Printf("UTC时间: %v\n", utc)
	// 获取时间戳(秒)
	unix := now.Unix()
	fmt.Printf("Unix时间戳: %d\n", unix)
	// 获取时间戳(纳秒)
	unixNano := now.UnixNano()
	fmt.Printf("Unix纳秒时间戳: %d\n", unixNano)
	year := now.Year()   // 年
	month := now.Month() // 月
	day := now.Day()     // 日
	// hour := now.Hour()       // 小时
	// minute := now.Minute()   // 分钟
	// second := now.Second()   // 秒
	// weekday := now.Weekday() // 星期几

	// 或者使用Date方法
	year, month, day = now.Date()
	fmt.Printf("%d年%d月%d日\n", year, month, day)

	// 格式化时间
	fmt.Println("格式化时间1", now.Format("2006-01-02 15:04:05")) // 2025-08-05 22:22:22
	fmt.Println("格式化时间2", now.Format("2006/01/02"))          // 2025/08/05
	fmt.Println("格式化时间3", now.Format("15:04:05"))            // 22:22:22

	// 解析时间字符串
	t, err := time.Parse("2006-01-02 15", "2025-03-15 22")
	if err != nil {
		panic(err)
	}
	fmt.Println(t) // 2025-03-15 00:00:00 +0000 UTC
	fmt.Println("-----------时间计算--------------")
	// 增加时间
	later := now.Add(time.Hour * 2) // 2小时后
	fmt.Println(later)
	// 减少时间
	earlier := now.Add(-time.Minute * 30) // 30分钟前
	fmt.Println(earlier)
	// 计算时间差
	duration := later.Sub(earlier)
	fmt.Printf("时间差: %v\n", duration) // 2h30m0s
	// 创建东八区（北京时间）的固定偏移时区
	// beijingLoc := time.FixedZone("CST", 8*60*60) // +8小时
	// 加载上海时区（等同于北京时间）
	shanghaiLoc, _ := time.LoadLocation("Asia/Shanghai")
	today1 := time.Date(2022, 10, 31, 0, 0, 0, 0, shanghaiLoc)
	fmt.Println("today1", today1)
	// AddDate 的特殊行为
	// 10月31日加1个月
	// time.Local指定本地时区
	today := time.Date(2022, 10, 31, 0, 0, 0, 0, time.Local)
	nextDay := today.AddDate(0, 1, 0)
	fmt.Println("AddDate", nextDay.Format("20060102")) // 输出: 20221201 (不是11月31日)
	t1 := time.Now()
	t2 := t1.Add(time.Hour)

	fmt.Println(t1.After(t2))  // false
	fmt.Println(t1.Before(t2)) // true
	fmt.Println(t1.Equal(t2))  // false
	fmt.Println("-----------定时任务--------------")
	// 创建1秒间隔的Ticker
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("执行任务:", t.Format("15:04:05"))
		}
	}()
	// 5秒后停止ticker
	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker已停止")
	// 方法二
	ticker1 := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case t := <-ticker1.C:
				fmt.Println("执行任务二:", t.Format("15:04:05"))
			case <-done:
				fmt.Println("收到停止信号二")
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	done <- true
	defer ticker1.Stop()
}
