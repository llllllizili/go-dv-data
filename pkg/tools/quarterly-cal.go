package tools

import (
	"time"
)

func GetCurrentQuarterDates() (int, string, string) {
	// 获取当前日期
	now := time.Now()
	year, month, _ := now.Date()
	// 计算当前月份所属的季度
	quarter := (int(month)-1)/3 + 1
	// 计算季度的第一天
	firstDay := time.Date(year, time.Month((quarter-1)*3+1), 1, 0, 0, 0, 0, time.UTC)
	// 计算季度的最后一天
	lastDay := firstDay.AddDate(0, 3, -1)

	firstDayStr := firstDay.Format("2006-01-02")
	lastDayStr := lastDay.Format("2006-01-02")

	return quarter, firstDayStr, lastDayStr

}
