package oodle_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/MR-AliHaashemi/uego/compression/oodle"
)

func TestDecompress(t *testing.T) {
	compressed, err := os.Open("../compressed.data")
	if err != nil {
		t.Fatal(err)
	}

	decompressed, err := os.Open("../decompressed.data")
	if err != nil {
		t.Fatal(err)
	}

	compressedBytes, err := io.ReadAll(compressed)
	if err != nil {
		t.Fatal(err)
	}

	decompressedBytes, err := io.ReadAll(decompressed)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := oodle.Decompress(compressedBytes, 65536)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(decompressedBytes, resp) {
		t.Errorf("invalid decompressed data")
	}
}
