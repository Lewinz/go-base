package timestamp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp(t *testing.T) {
	assert := assert.New(t)

	now := Now()
	assert.NotNil(now)

	uts := Unix(1490328882, 0)
	assert.NotNil(uts)
	assert.Equal(uts.Unix(), int64(1490328882))

	date := Date(2018, 11, 11, 0, 0, 0, 0, time.Local)
	assert.NotNil(date)
	assert.Equal(date.Year(), 2018)
	assert.Equal(date.Month(), time.Month(11))
	assert.Equal(date.Day(), 11)

	times := New()
	assert.NotNil(times)

	err := times.UnmarshalJSON([]byte("2018-11-15 17:56:17"))
	assert.Nil(err)
	assert.False(times.IsZero())

	err = times.UnmarshalJSON([]byte("1490328882"))
	assert.Nil(err)
	assert.False(times.IsZero())

	b, err := times.MarshalJSON()
	assert.Nil(err)
	assert.True(len(b) > 0)
	assert.Equal(string(b), "1490328882")

	err = times.Scan(time.Now())
	assert.Nil(err)
	assert.False(times.IsZero())

	value, err := times.Value()
	assert.Nil(err)
	assert.NotNil(value)

	_, ok := value.(time.Time)
	assert.True(ok)
}
