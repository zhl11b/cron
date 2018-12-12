package cron

import (
	"fmt"
	"time"

	"github.com/cxuhua/xweb/now"
)

type PeriodUnit int

const (
	_PeriodUnit PeriodUnit = iota
	UnitDay
	UnitWeek
	UnitMonth
	UnitYear
)

type ScheduleEvery struct {
	StartTime time.Time
	Period    int
	Unit      PeriodUnit
}

// NewScheduleEvery 新建时间周期
func NewScheduleEvery(period int, unit PeriodUnit, startTime time.Time) *ScheduleEvery {
	return &ScheduleEvery{
		Period:    period,
		Unit:      unit,
		StartTime: startTime,
	}
}

// Next 下一个时间
func (s *ScheduleEvery) Next(t time.Time) time.Time {
	now.FirstDayMonday = true
	switch s.Unit {
	case UnitDay:
		count := -int(now.New(s.StartTime).BeginningOfDay().Sub(t).Hours()) / (24 * s.Period)
		fmt.Println(count)
		return now.New(s.StartTime).BeginningOfDay().Add(time.Duration(count+1) * time.Hour * 24 * time.Duration(s.Period))
	case UnitWeek:
		count := -int(now.New(s.StartTime).BeginningOfWeek().Sub(t).Hours()) / (24 * 7 * s.Period)
		fmt.Println(count)
		return now.New(s.StartTime).BeginningOfWeek().Add(time.Duration(count+1) * time.Hour * 24 * 7 * time.Duration(s.Period))
	case UnitMonth:
		temp := now.New(s.StartTime).BeginningOfMonth()
		for i := 0; temp.Before(t); i++ {
			temp = temp.AddDate(0, s.Period, 0)
		}
		return temp
	case UnitYear:
		temp := now.New(s.StartTime).BeginningOfYear()
		for i := 0; temp.Before(t); i++ {
			temp = temp.AddDate(s.Period, 0, 0)
		}
		return temp
	default:
		return time.Time{}
	}
}

// Prev 上一个时间
func (s *ScheduleEvery) Prev(t time.Time) time.Time {
	switch s.Unit {
	case UnitDay:
		count := -int(now.New(s.StartTime).BeginningOfDay().Sub(t).Hours()) / (24 * s.Period)
		return now.New(s.StartTime).BeginningOfDay().Add(time.Duration(count) * time.Hour * 24 * time.Duration(s.Period))
	case UnitWeek:
		count := -int(now.New(s.StartTime).BeginningOfWeek().Sub(t).Hours()) / (24 * 7 * s.Period)
		return now.New(s.StartTime).BeginningOfWeek().Add(time.Duration(count) * time.Hour * 24 * 7 * time.Duration(s.Period))
	case UnitMonth:
		temp := now.New(s.StartTime).BeginningOfMonth()
		for i := 0; temp.Before(t); i++ {
			temp = temp.AddDate(0, s.Period, 0)
		}
		return temp.AddDate(0, -s.Period, 0) // 少一个周期
	case UnitYear:
		temp := now.New(s.StartTime).BeginningOfYear()
		for i := 0; temp.Before(t); i++ {
			temp = temp.AddDate(s.Period, 0, 0)
		}
		return temp.AddDate(-s.Period, 0, 0) // 少一个周期
	default:
		return time.Time{}
	}
}
