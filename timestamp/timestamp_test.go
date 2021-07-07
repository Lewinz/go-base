package timestamp_test

import (
	"testing"
	"time"

	"github.com/Lewinz/go-base/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestTimestamp(t *testing.T) {
	assert := assert.New(t)

	now := timestamp.Now()
	assert.NotNil(now)

	uts := timestamp.Unix(1490328882, 0)
	assert.NotNil(uts)
	assert.Equal(uts.Unix(), int64(1490328882))

	date := timestamp.Date(2018, 11, 11, 0, 0, 0, 0, time.Local)
	assert.NotNil(date)
	assert.Equal(date.Year(), 2018)
	assert.Equal(date.Month(), time.Month(11))
	assert.Equal(date.Day(), 11)

	timestamp := timestamp.New()
	assert.NotNil(timestamp)

	err := timestamp.UnmarshalJSON([]byte("2018-11-15 17:56:17"))
	assert.Nil(err)
	assert.False(timestamp.IsZero())

	err = timestamp.UnmarshalJSON([]byte("1490328882"))
	assert.Nil(err)
	assert.False(timestamp.IsZero())

	b, err := timestamp.MarshalJSON()
	assert.Nil(err)
	assert.True(len(b) > 0)
	assert.Equal(string(b), "1490328882")

	err = timestamp.Scan(time.Now())
	assert.Nil(err)
	assert.False(timestamp.IsZero())

	value, err := timestamp.Value()
	assert.Nil(err)
	assert.NotNil(value)

	_, ok := value.(time.Time)
	assert.True(ok)
}
