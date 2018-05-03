package util

import (
	"time"
)

/*
	go中的日期格式为2006-01-02 15:04:05
*/

type DatePair struct {
	BeginDate time.Time
	EndDate   time.Time
}

//获取两个日期之间的工作日
func GetWeeks(t1 time.Time, t2 time.Time) []DatePair {
	dateList := []DatePair{}
	t0 := t1
	for ; t1.Before(t2); t1 = t1.Add(time.Hour * 24) {
		if t1.Weekday() == time.Saturday || t1.Weekday() == time.Sunday {
			t0 = t1.Add(time.Hour * 24)
			continue
		}
		if t1.Weekday() == time.Friday {
			datePair := DatePair{t0, t1}
			dateList = append(dateList, datePair)
			t0 = t1.Add(time.Hour * 24 * 3)
		}
	}
	if t2.Weekday() != time.Saturday && t2.Weekday() != time.Sunday {
		datePair := DatePair{t0, t2}
		dateList = append(dateList, datePair)
	}
	return dateList
}

//运用补码的思路计算起始日不是周日的周起始日期
func GetSpecialWeeks(pattern string, t1 time.Time, weekBegin time.Weekday) (datePair DatePair) {
	datePair = DatePair{}
	t2 := t1
	i := 0
	if t1.Weekday() > weekBegin {
		i = int(weekBegin - t1.Weekday())
	} else {
		i = int(t1.Weekday() - weekBegin + 7)
	}

	for j := 0; j <= i; j++ {
		if t2.Weekday() == weekBegin {
			break
		}
		t2 = t2.Add(time.Hour * -1 * 24)

	}
	datePair.BeginDate = t2
	datePair.EndDate = t1
	return datePair
}

//格式化日期
func GetDateString(pattern string) (date string) {
	now := time.Now()
	return now.Format(pattern)
}

//将日期字符串转换成日期值
func ConverStrToDate(pattern string, inputDate string) (tm time.Time, err error) {
	tm, err = time.Parse(pattern, inputDate)
	return tm, err
}

//转换日期字符串格式
func FormatDateString(pattern string, inputDate string, newPattern string) (dateStr string, err error) {
	tm, err := ConverStrToDate(pattern, inputDate)
	if err != nil {
		return "", err
	} else {
		return tm.Format(newPattern), nil
	}

}

//加一天，比较月份，如果不一样就是本月最后一天
func IsLastDayOfMonth(pattern string, inputDate string) bool {
	tm, err := ConverStrToDate(pattern, inputDate)
	if err != nil {
		return false
	} else {
		tm1 := tm.Add(time.Hour * 24)
		return tm.Month() == tm1.Month()
	}
}

//加一天，比较年份，如果不一样就是本年最后一天
func IsLastDayOfYear(pattern string, inputDate string) bool {
	tm, err := ConverStrToDate(pattern, inputDate)
	if err != nil {
		return false
	} else {
		tm1 := tm.Add(time.Hour * 24)
		return tm.Year() == tm1.Year()
	}
}
