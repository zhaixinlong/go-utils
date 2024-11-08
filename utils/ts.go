package util

import (
	"fmt"
	"math"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	TsExpNever       = "2099-01-01 00:00:00"
	formatNormal     = time.RFC3339
	formatShort      = "2006-01-02 15:04:05"
	formatCompact    = "20060102150405"
	formatDay        = "2006-01-02"
	formatCompactDay = "20060102"
)

func FormatTs(t time.Time) string {
	return t.Format(formatNormal)
}

func FormatShortTs(t time.Time) string {
	return t.Format(formatShort)
}

func FormatShortStartOfDay(t time.Time) string {
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return FormatShortTs(startOfDay)
}

func FormatShortEndOfDay(t time.Time) string {
	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return FormatShortTs(endOfDay)
}

func FormatCompactTs(t time.Time) string {
	return t.Format(formatCompact)
}

func FormatCompactTsNano(t time.Time) string {
	return fmt.Sprintf("%s_%09d", t.Format(formatCompact), t.Nanosecond())
}

func GetShortTsTime(ts string) (time.Time, error) {
	return time.ParseInLocation(formatShort, ts, time.Local)
}

func FormatTsToShortLocalTime(ts string) *time.Time {
	t, err := GetShortTsTime(ts)
	if err != nil {
		panic(err)
	}
	return &t
}

func FormatTsToLocalTime(ts string) *time.Time {
	t, err := time.ParseInLocation(formatNormal, ts, time.Local)
	if err != nil {
		logrus.Errorf("FormatTsToLocalTime %s error:%#v", ts, err)
		return nil
	}
	return &t
}

func FormatTimeToLocal(ts string) string {
	t := FormatTsToLocalTime(ts)
	if t == nil {
		return ""
	}
	return FormatTs(*t)
}

func GetShortTsDuration(ts1, ts2 string) (time.Duration, error) {
	t1, err := GetShortTsTime(ts1)
	if err != nil {
		return 0, err
	}
	t2, err := GetShortTsTime(ts2)
	if err != nil {
		return 0, err
	}
	return t1.Sub(t2), nil
}

func FormatTsNil(t *time.Time) string {
	if t == nil {
		return ""
	}
	return FormatShortTs(*t)
}

// 时间戳转日期
func ConvertTime(utime uint64) string {
	format := time.Unix(int64(utime), 0).Format("2006-01-02 00:00:00")
	return format
}

// 时间戳转天数
func Days(timestampFrom, timestampTo int64) int {
	var midnightUnix = func(t time.Time) int64 {
		y, m, d := t.Date()
		return time.Date(y, m, d+1, 0, 0, 0, 0, time.Local).Unix()
	}
	var days = 0
	for {
		if midnightUnix(time.Unix(timestampFrom, 0).AddDate(0, 0, days)) >= timestampTo {
			days++
			break
		}
		days++
	}
	return days
}

func ParseTime(timeStr string) (*time.Time, error) {
	ts, err := time.ParseInLocation(formatShort, timeStr, time.Local)
	if err != nil {
		logrus.Errorf("ParseTime %s error:%#v", timeStr, err)
		return nil, err
	}
	return &ts, nil
}

func FormatTimestampToShortTs(pbTime *timestamppb.Timestamp) string {
	return FormatShortTs(pbTime.AsTime().Local())
}

func FormatTimestampToTime(pbTime *timestamppb.Timestamp) time.Time {
	return pbTime.AsTime().Local()
}

func FormatShortTsToTimestamp(shortTsTime string) *timestamppb.Timestamp {
	t, err := GetShortTsTime(shortTsTime)
	if err != nil {
		return timestamppb.Now()
	}
	return timestamppb.New(t)
}

func GetDayFormatTime(t time.Time) time.Time {
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return startOfDay
}

func TimeStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func TimeStartOfHour(t time.Time, hour int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, t.Location())
}

func TimeAddDay(t time.Time, day int32) time.Time {
	return t.Add(time.Duration(day) * 24 * time.Hour)
}

func GetDiffDays(startTime time.Time, endTime time.Time) int32 {
	dura := endTime.Sub(startTime).Hours()
	return int32(math.Abs(math.Floor(dura / 24)))
}

func GetsDaysCharacter(t time.Time, day int) string {
	return FormatDay(TimeAddDay(t, int32(day)))
}

func FormatDay(t time.Time) string {
	return t.Format(formatDay)
}

func VerifyTimeIsIntersectForStr(bt1, et1, bt2, et2 string) (bool, error) {
	time1Start, err := GetShortTsTime(bt1)
	if err != nil {
		return false, err
	}
	time1End, err := GetShortTsTime(et1)
	if err != nil {
		return false, err
	}

	time2Start, err := GetShortTsTime(bt2)
	if err != nil {
		return false, err
	}
	time2End, err := GetShortTsTime(et2)
	if err != nil {
		return false, err
	}
	return VerifyTimeIsIntersectForTime(time1Start, time1End, time2Start, time2End)
}

func VerifyTimeIsIntersectForTime(time1Start, time1End, time2Start, time2End time.Time) (bool, error) {
	if time1End.Before(time1Start) {
		return false, nil
	}
	if time2End.Before(time2Start) {
		return false, nil
	}

	// left
	if time2Start.After(time1End) || time2Start.Equal(time1End) {
		return true, nil
	}
	// right
	if time1Start.After(time2End) || time1Start.Equal(time2End) {
		return true, nil
	}
	return false, nil
}

func GetCompactTsTime(ts string) (time.Time, error) {
	return time.ParseInLocation(formatCompact, ts, time.Local)
}

func FormatCompactDay(t time.Time) string {
	return t.Format(formatCompactDay)
}

func GetFormatDay(ts string) (time.Time, error) {
	return time.ParseInLocation(formatDay, ts, time.Local)
}
