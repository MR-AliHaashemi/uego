// Package uereader contains some UnrealEngine's specific
// reader methods
package uereader

import (
	"encoding/binary"
	"errors"
	"io"
)

var (
	ErrFStringNotNullTerminated = errors.New("string is not null terminated")
)

func Read[T any](r io.Reader) (data T, err error) {
	err = binary.Read(r, binary.LittleEndian, &data)
	return
}

func ReadArray[T any](r io.Reader, n uint64) ([]T, error) {
	var out []T

	for i := uint64(0); i < n; i++ {
		var val T

		if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
			return nil, err
		}

		out = append(out, val)
	}

	return out, nil
}

func ReadArrayN[T any](r io.Reader, n int64, fn func(io.Reader) (T, error)) ([]T, error) {
	var out []T

	for i := int64(0); i < n; i++ {
		val, err := fn(r)
		if err != nil {
			return nil, err
		}

		out = append(out, val)
	}

	return out, nil
}

// FString reads a null-terminated string starting with the length from r
func FString(r io.Reader) (string, error) {
	length, err := Read[uint32](r)
	if err != nil || length == 0 {
		return "", err
	}

	buf := make([]byte, length)
	if err = binary.Read(r, binary.LittleEndian, &buf); err != nil {
		return "", err
	}

	if buf[len(buf)-1] != 0x0 { // ensure it's null-terminated
		return "", ErrFStringNotNullTerminated
	}

	return string(buf[:len(buf)-1]), nil
}

func TArray[T any](r io.Reader, fn func(io.Reader) (T, error)) (out []T, err error) {
	length, err := Read[uint32](r)
	if err != nil || length == 0 {
		return nil, err
	}

	for i := uint32(0); i < length; i++ {
		data, err := fn(r)
		if err != nil {
			return nil, err
		}

		out = append(out, data)
	}

	return out, nil
}
