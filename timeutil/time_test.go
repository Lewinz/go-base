package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 测试时间操作工具
func TestTimeUtils(t *testing.T) {
	// 按时区时间格式转换time
	time1, err := ParseTimeToLocal("2021-06-13T00:00:00Z", DataTimeFormatSecondUTCLocation)
	localTimeStr1 := time1.Format(DataTimeFormatSecond)
	assert.Nil(t, err)
	assert.True(t, localTimeStr1 == "2021-06-13 08:00:00")

	// 将 其他时区时间转化为 Local 时间
	localTime, err := ParseTimeToLocal("2021-06-13 00:00:00", DataTimeFormatSecond, time.UTC)
	localTimeStr := localTime.Format(DataTimeFormatSecond)
	assert.Nil(t, err)
	assert.True(t, localTimeStr == "2021-06-13 08:00:00")

	// 空值单测
	nilTestTime, err := ParseTimeToLocal("", DataTimeFormatSecond, time.UTC)
	assert.Nil(t, err)

	// 将 当前服务器时间转化为 UTC 时间
	UTCTime, err := ParseTimeToUTC("2021-06-13 08:00:00", DataTimeFormatSecond)
	UTCTimeStr := UTCTime.Format(DataTimeFormatSecond)
	assert.Nil(t, err)
	assert.True(t, UTCTimeStr == "2021-06-13 00:00:00")

	// 空值单测
	_, err = ParseTimeToUTC("", DataTimeFormatSecond)
	assert.Nil(t, err)

	// 将 时间转换为时间戳
	unix, err := ParseUnix("2021-06-13 08:00:00", DataTimeFormatSecond, MilliSecondUnix, time.UTC)
	assert.Nil(t, err)
	assert.True(t, unix == 1623571200000)

	// 空值单测
	_, err = ParseUnix("", DataTimeFormatSecond, MilliSecondUnix, time.UTC)
	assert.Nil(t, err)

	// 将 时间转换为时间戳
	unix2, err := ParseTimeUnix(localTime, SecondUnix)
	assert.Nil(t, err)
	assert.True(t, unix2 == 1623542400)

	// 空值单测
	_, err = ParseTimeUnix(nilTestTime, MilliSecondUnix)
	assert.Nil(t, err)
}
