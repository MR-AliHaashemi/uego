package ue

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/google/uuid"
)

var (
	ErrFStringNotNullTerminated = errors.New("string is not null terminated")
)

type FArchive = *Archive

type Archive struct{ io.ReadSeeker }

// Size returns the archive bytes size
func (r *Archive) Size() (int64, error) {
	curr, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	size, err := r.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}

	_, err = r.Seek(curr, io.SeekStart)
	if err != nil {
		return 0, err
	}

	return size, nil
}

// Position returns the current reader's position
func (r *Archive) Position() (int64, error) {
	return r.Seek(0, io.SeekCurrent)
}

func (r *Archive) Bool() (bool, error) {
	data, err := r.Byte()
	if err != nil {
		return false, err
	}
	return data != 0x0, nil
}

func (r *Archive) FBool() (bool, error) {
	// UnrealEngine stores boolean as uint8 for whatever reason
	data, err := r.UInt8()
	if err != nil {
		return false, err
	}

	return data != 0, nil
}

func (r *Archive) Byte() (byte, error) {
	data := make([]byte, 1)
	_, err := r.Read(data)
	if err != nil {
		return 0, err
	}

	return data[0], nil
}

func (r *Archive) Bytes(size int) ([]byte, error) {
	data := make([]byte, size)
	if _, err := r.Read(data); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Archive) UInt8() (uint8, error) {
	return r.Byte()
}

func (r *Archive) UInt16() (uint16, error) {
	data := make([]byte, 2)
	_, err := r.Read(data[:])
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(data), nil
}

func (r *Archive) UInt32() (uint32, error) {
	data := make([]byte, 4)
	_, err := r.Read(data[:])
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(data), nil
}

func (r *Archive) UInt64() (uint64, error) {
	data := make([]byte, 8)
	_, err := r.Read(data[:])
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(data), nil
}

func (r *Archive) Int8() (int8, error) {
	num, err := r.UInt8()
	return int8(num), err
}

func (r *Archive) Int16() (int16, error) {
	num, err := r.UInt16()
	return int16(num), err
}

func (r *Archive) Int32() (int32, error) {
	num, err := r.UInt32()
	return int32(num), err
}

func (r *Archive) Int64() (int64, error) {
	num, err := r.UInt64()
	return int64(num), err
}

func (r *Archive) UUID() (uuid.UUID, error) {
	var data uuid.UUID
	_, err := r.Read(data[:])
	if err != nil {
		return uuid.Nil, err
	}

	return data, nil
}

func (r *Archive) BigEndianUUID() (guid uuid.UUID, err error) {
	data := make([]uint32, 4)
	if err = binary.Read(r, binary.BigEndian, &data); err != nil {
		return uuid.Nil, err
	}

	for i, v := range data {
		binary.LittleEndian.PutUint32(guid[i*4:(i+1)*4], v)
	}

	return guid, nil
}

func (r *Archive) FString() (string, error) {
	size, err := r.UInt32()
	if err != nil || size == 0 {
		return "", err
	}

	data := make([]byte, size)
	if _, err = r.Read(data); err != nil {
		return "", err
	}

	// ensure it's null-terminated
	if data[len(data)-1] != 0x0 {
		return "", ErrFStringNotNullTerminated
	}

	return string(data[:len(data)-1]), nil // avoid the null character while returning
}

func (r *Archive) ShaHash() (SHAHash, error) {
	var data SHAHash
	_, err := r.Read(data[:])

	return data, err
}

func NewArchive(r io.ReadSeeker) *Archive { return &Archive{r} }
