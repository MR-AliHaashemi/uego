package ue

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"

	"github.com/er-azh/goodle"
)

var ErrUnknownCompressionMethod = errors.New("unknown compression method")

type CompressionMethod string

const (
	CompressionMethodNone    CompressionMethod = ""
	CompressionMethodZlib    CompressionMethod = "Zlib"
	CompressionMethodGzip    CompressionMethod = "Gzip"
	CompressionMethodOodle   CompressionMethod = "Oodle"
	CompressionMethodUnknown CompressionMethod = "UNKNOWN"
)

func GetCompressionMethod(n int) CompressionMethod {
	switch n {
	case 1:
		return CompressionMethodZlib
	case 2:
		return CompressionMethodGzip
	case 3:
		return CompressionMethodOodle
	}

	return CompressionMethodNone
}

func GetCompressionByString(s string) CompressionMethod {
	switch s {
	case "":
		return CompressionMethodNone
	case "Zlib":
		return CompressionMethodZlib
	case "Gzip":
		return CompressionMethodGzip
	case "Oodle":
		return CompressionMethodOodle
	}

	return CompressionMethodUnknown
}

func Decompress(compressed []byte, compressedOffset, decompressedSize int, method CompressionMethod) (io.ReadCloser, error) {
	switch method {
	case CompressionMethodNone:
		return io.NopCloser(bytes.NewReader(compressed)), nil
	case CompressionMethodZlib:
		return zlib.NewReader(bytes.NewReader(compressed))
	case CompressionMethodGzip:
		return gzip.NewReader(bytes.NewReader(compressed))
	case CompressionMethodOodle:
		data, err := goodle.Decompress(compressed, decompressedSize)
		if err != nil {
			return nil, err
		}
		return io.NopCloser(bytes.NewReader(data)), nil
	}

	return nil, ErrUnknownCompressionMethod
}
