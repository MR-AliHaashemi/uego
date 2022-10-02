//go:build windows
// +build windows

package oodle

import (
	"syscall"
	"unsafe"
)

var (
	Oo2CoreDLL         = syscall.NewLazyDLL("oo2core_9_win64.dll")
	OodleLZ_Decompress = Oo2CoreDLL.NewProc("OodleLZ_Decompress")
)

// Decompress returns raw (decompressed) len received;
// returns 0 (ErrOodleZFailed) if it detects corruption.
//
// https://github.com/EpicGames/UnrealEngine/blob/46544fa5e0aa9e6740c19b44b0628b72e7bbd5ce/Engine/Source/Programs/Horde/HordeStorage/Jupiter.Common/Utils/OodleCompressor.cs#L126
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
	r1, _, _ := OodleLZ_Decompress.Call(
		uintptr(compBuf),
		uintptr(compBufSize),
		uintptr(rawBuf),
		uintptr(rawLen),
		uintptr(fuzzSafe),
		uintptr(checkCRC),
		uintptr(verbosity),
		uintptr(decBufBase),
		uintptr(decBufSize),
		uintptr(fpCallback),
		uintptr(callbackUserData),
		uintptr(decoderMemory),
		uintptr(decoderMemorySize),
		uintptr(threadPhase),
	)

	return int64(r1)
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
