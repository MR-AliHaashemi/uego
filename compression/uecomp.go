package compression

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"

	"github.com/MR-AliHaashemi/uego/compression/oodle"
)

var ErrUnknownMethod = errors.New("unknown compression method")

type Method string

const (
	MethodNone    Method = ""
	MethodZlib    Method = "Zlib"
	MethodGzip    Method = "Gzip"
	MethodOodle   Method = "Oodle"
	MethodUnknown Method = "UNKNOWN"
)

func GetMethod(n int) Method {
	switch n {
	case 1:
		return MethodZlib
	case 2:
		return MethodGzip
	case 3:
		return MethodOodle
	}

	return MethodNone
}

func GetStringMethod(s string) Method {
	switch s {
	case "":
		return MethodNone
	case "Zlib":
		return MethodZlib
	case "Gzip":
		return MethodGzip
	case "Oodle":
		return MethodOodle
	}

	return MethodUnknown
}

func Decompress(compressed []byte, compressedOffset, decompressedSize int, method Method) (io.ReadCloser, error) {
	switch method {
	case MethodNone:
		return io.NopCloser(bytes.NewReader(compressed)), nil
	case MethodZlib:
		return zlib.NewReader(bytes.NewReader(compressed))
	case MethodGzip:
		return gzip.NewReader(bytes.NewReader(compressed))
	case MethodOodle:
		data, err := oodle.Decompress(compressed, decompressedSize)
		if err != nil {
			return nil, err
		}
		return io.NopCloser(bytes.NewReader(data)), nil
	}

	return nil, ErrUnknownMethod
}
