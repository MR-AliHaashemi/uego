//go:build linux
// +build linux

package oodle

//#cgo linux LDFLAGS: -loo2corelinux64_dbg -lm -lstdc++ -L .
//#include "./include/oodle2.h"
import "C"

import "unsafe"

func DecompressRaw(
	compBuf unsafe.Pointer,
	compBufSize int64,
	rawBuf unsafe.Pointer,
	rawLen int64,
	fuzzSafe FuzzSafe,
	checkCRC CheckCRC,
	verbosity Verbosity,
	decBufBase unsafe.Pointer,
	decBufSize int64,
	fpCallback int64,
	callbackUserData int64,
	decoderMemory unsafe.Pointer,
	decoderMemorySize int64,
	threadPhase DecodeThreadPhase,
) int64 {
	return int64(C.OodleLZ_Decompress(
		compBuf,
		C.long(compBufSize),
		rawBuf,
		C.long(rawLen),
		C.OodleLZ_FuzzSafe(fuzzSafe),
		C.OodleLZ_CheckCRC(checkCRC),
		C.OodleLZ_Verbosity(verbosity),
		decBufBase,
		C.long(decBufSize),
		(*C.OodleDecompressCallback)(fpCallback),
		callbackUserData,
		decoderMemory,
		C.long(decoderMemorySize),
		C.OodleLZ_Decode_ThreadPhase(threadPhase),
	))
}

func Decompress(compressed []byte, decompressedSize int) ([]byte, error) {
	rawBuf := make([]byte, decompressedSize)

	if DecompressRaw(
		unsafe.Pointer(&compressed[0]),
		int64(len(compressed)),
		unsafe.Pointer(&rawBuf[0]),
		int64(len(rawBuf)),

		FuzzSafeYes,      // fuzzSafe
		CheckCRCNo,       // checkCRC
		VerbosityNone,    // verbosity
		nil,              // decBufBase
		0,                // decBufSize
		0,                // fpCallback
		0,                // callbackUserData
		nil,              // decoderMemory
		0,                // decoderMemorySize
		DecodeUnthreaded, // threadPhase
	) == 0 {
		return nil, ErrOodleLZFailed
	}

	return rawBuf, nil
}
