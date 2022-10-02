package compression

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"

	"github.com/MR-AliHaashemi/uego/compression/oodle"
)

type Methods string

const (
	MethodNone  Methods = ""
	MethodZlib  Methods = "Zlib"
	MethodGzip  Methods = "Gzip"
	MethodOodle Methods = "Oodle"
)

func Decompress(compressed []byte, compressedOffset, decompressedSize int, method Methods) (io.ReadCloser, error) {
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
	default:
		return nil, fmt.Errorf("unknown compression method %v", method)
	}
}
