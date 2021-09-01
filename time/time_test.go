package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 测试时间操作工具
func TestTimeUtils(t *testing.T) {
	// 按时区时间格式转换time
	_, err := ParseTimeToLocal("2021-06-13T00:00:00Z", DataTimeFormatSecondUTCLocation)
	assert.Nil(t, err)

	_, err = ParseTimeToLocal("2021-06-13 08:00:00", DataTimeFormatSecond)
	assert.Nil(t, err)

	// 将 时间转换为时间戳
	_, err = ParseUnix("2021-06-13T00:00:00Z", DataTimeFormatSecondUTCLocation, NanoSecondUnix, time.UTC)
	assert.Nil(t, err)

	// 空值单测
	nilTestTime, err := ParseTimeToLocal("", DataTimeFormatSecond, time.UTC)
	assert.Nil(t, err)

	// 空值单测
	_, err = ParseTimeToUTC("", DataTimeFormatSecond)
	assert.Nil(t, err)

	// 空值单测
	_, err = ParseUnix("", DataTimeFormatSecond, MilliSecondUnix, time.UTC)
	assert.Nil(t, err)

	// 空值单测
	_, err = ParseTimeUnix(nilTestTime, MilliSecondUnix)
	assert.Nil(t, err)
}
