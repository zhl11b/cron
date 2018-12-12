package cron

import (
	"testing"
	"time"
)

func TestScheduleEvery_Next(t *testing.T) {
	start, err := time.Parse("2006-01-02 15:04:05", "2018-12-08 16:05:02")
	if err != nil {
		t.Error(err)
	}
	s1 := NewScheduleEvery(3, UnitDay, start)
	t.Logf("day:%v prev:%v\n", s1.Next(time.Now()), s1.Prev(time.Now()))
	s2 := NewScheduleEvery(3, UnitWeek, start)
	t.Logf("week:%v prev:%+v\n", s2.Next(time.Now()), s2.Prev(time.Now()))
	s3 := NewScheduleEvery(3, UnitMonth, start)
	t.Logf("month:%v prev:%v\n", s3.Next(time.Now()), s3.Prev(time.Now()))
	s4 := NewScheduleEvery(3, UnitYear, start)
	t.Logf("year:%v prev:%v\n", s4.Next(time.Now()), s4.Prev(time.Now()))

}
