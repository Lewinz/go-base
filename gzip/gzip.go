package gzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// Compression 字符串压缩
func Compression(str string) ([]byte, error) {
	var buffer bytes.Buffer
	gz := gzip.NewWriter(&buffer)

	if _, err := gz.Write([]byte(str)); err != nil {
		return nil, err
	}

	// 不能放在 defer 里
	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Decompression 字符串解压缩
func Decompression(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}
