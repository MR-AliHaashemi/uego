package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"strings"
)

var (
	ErrAESKeyWrongLength       = errors.New("aes key must be 32 bytes long")
	ErrAESKeyStringWrongLength = errors.New("aes key must be 64 characters long without starting 0x")
)

type AES struct {
	block cipher.Block
	key   []byte
}

func NewAES(key []byte) (*AES, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AES{
		block: block,
		key:   key,
	}, nil
}

func NewAESString(key string) (*AES, error) {
	key = strings.TrimPrefix(key, "0x")

	if len(key) != 64 {
		return nil, ErrAESKeyStringWrongLength
	}

	rawKey, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(rawKey)
	if err != nil {
		return nil, err
	}

	return &AES{
		block: block,
		key:   rawKey,
	}, nil
}

func (a *AES) Decrypt(data []byte) []byte {
	decrypted := make([]byte, len(data))

	for bs, be := 0, aes.BlockSize; bs < len(data); bs, be = bs+aes.BlockSize, be+aes.BlockSize {
		a.block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func (a *AES) Key() []byte {
	return a.key
}

func (a *AES) KeyString() string {
	return hex.EncodeToString(a.key)
}
