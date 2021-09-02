package gzip

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCompression(t *testing.T) {
	str := "gzip compression test"
	comStr, err := Compression(str)
	assert.Nil(t, err)
	assert.NotEmpty(t, comStr)

	decomByte, err := Decompression(comStr)
	assert.Nil(t, err)
	assert.NotEmpty(t, decomByte)
	assert.True(t, string(decomByte) == str)

	str2 := ""
	comStr2, err := Compression(str2)
	assert.Nil(t, err)

	decomByte2, err := Decompression(comStr2)
	assert.Nil(t, err)
	assert.True(t, string(decomByte2) == str2)
}
