package time

import (
	"fmt"
	"time"
)

const (
	// DataFormat day
	DataFormat = "2006-01-02"
	// DataTimeFormat minute
	DataTimeFormat = "2006-01-02 15:0"
	// DataTimeFormatSecond second
	DataTimeFormatSecond = "2006-01-02 15:04:05"
	// DataTimeFormatSecond2 second
	DataTimeFormatSecond2 = "2006-01-02T15:04:05"
	// DataTimeFormatSecondUTCLocation second location UTC
	DataTimeFormatSecondUTCLocation = "2006-01-02T15:04:05Z"
	// DataTimeFormatMilliSecond milli Second location
	DataTimeFormatMilliSecond = "2006-01-02T15:04:05.000000"
)

// UnixType 时间戳精度
type UnixType string

// 时间戳精度 Enum
const (
	SecondUnix      UnixType = "SecondUnix"      // 秒
	MilliSecondUnix UnixType = "MilliSecondUnix" // 毫秒
	NanoSecondUnix  UnixType = "NanoSecondUnix"  // 纳秒
)

// ParseTime 将时间字符串按 layout 格式转换为时间格式
func ParseTime(t, layout string, loc ...*time.Location) (tt time.Time, err error) {
	if t == "" || layout == "" {
		return time.Time{}, nil
	}

	location := time.Local
	if len(loc) > 0 {
		location = loc[0]
	}

	switch layout {
	// 带时区格式，使用格式中的时区指定转换
	case DataTimeFormatSecondUTCLocation:
		return time.Parse(layout, t)
	default:
		return time.ParseInLocation(layout, t, location)
	}
}

// ParseTimeToLocal 将时区时间转换为当前服务器时间，例如 UTC 时区 2021-06-13T00:00:00Z 将转化为 2021-06-13T08:00:00Z
func ParseTimeToLocal(t, layout string, loc ...*time.Location) (localTime time.Time, err error) {
	tt, err := ParseTime(t, layout, loc...)
	if err != nil {
		return
	}

	localTime = tt.In(time.Local)
	return
}

// ParseTimeToUTC 将时区时间转换为UTC时区时间
func ParseTimeToUTC(t, layout string, loc ...*time.Location) (utcTime time.Time, err error) {
	tt, err := ParseTime(t, layout, loc...)
	if err != nil {
		return
	}

	utcTime = tt.In(time.UTC)
	return
}

// ParseUnix 转换为时间戳
func ParseUnix(t, layout string, unixType UnixType, loc ...*time.Location) (tInt64 int64, err error) {
	time, err := ParseTime(t, layout, loc...)
	if err != nil {
		return
	}

	tInt64, err = ParseTimeUnix(time, unixType)
	return
}

// ParseTimeUnix 转换为时间戳
func ParseTimeUnix(time time.Time, unixType UnixType) (tInt64 int64, err error) {
	switch unixType {
	case SecondUnix:
		return time.Unix(), nil
	case MilliSecondUnix:
		return time.UnixNano() / 1e6, nil
	case NanoSecondUnix:
		return time.UnixNano(), nil
	default:
		return 0, fmt.Errorf("time parse unix err")
	}
}

// ParseTimeInt64 将时间字符串按 layout 格式转换为时间戳格式，毫秒
func ParseTimeInt64(t, layout string, loc ...*time.Location) (tInt64 int64, err error) {
	tt, err := ParseTime(t, layout, loc...)
	if err != nil {
		return
	}

	tInt64 = tt.Unix() * 1000
	return
}

// ParseUTCMillSecond 将 UTC 字符串转换成 local time 毫秒
func ParseUTCMillSecond(t, layout string) (tInt64 int64, err error) {
	tt, err := ParseTimeToLocal(t, layout, time.UTC)
	if err != nil {
		return
	}

	return tt.Unix() * 1000, nil
}
