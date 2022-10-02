package oodle

import "errors"

var (
	ErrOodleLZFailed = errors.New("decompression failed")
)

// https://github.com/EpicGames/UnrealEngine/blob/46544fa5e0aa9e6740c19b44b0628b72e7bbd5ce/Engine/Source/Programs/Horde/HordeStorage/Jupiter.Common/Utils/OodleCompressor.cs#L10
type Compressor int32

const (
	CompressorInvalid Compressor = -1
	CompressorNone    Compressor = 3 // None = memcpy  pass through uncompressed bytes

	CompressorKraken    Compressor = 8  // Fast decompression and high compression ratios  amazing!
	CompressorLeviathan Compressor = 13 // Leviathan = Kraken's big brother with higher compression  slightly slower decompression.
	CompressorMermaid   Compressor = 9  // Mermaid is between Kraken & Selkie - crazy fast  still decent compression.
	CompressorSelkie    Compressor = 11 // Selkie is a super-fast relative of Mermaid.  For maximum decode speed.
	CompressorHydra     Compressor = 12 // Hydra the many-headed beast = Leviathan  Kraken  Mermaid  or Selkie (see $ AboutHydra)

	/* Deprecated compressors
	CompressorBitKnit Compressor = 10 // no longer supported as of Oodle 2.9.0
	CompressorLZB16   Compressor = 4  // DEPRECATED but still supported
	CompressorLZNA    Compressor = 7  // no longer supported as of Oodle 2.9.0
	CompressorLZH     Compressor = 0  // no longer supported as of Oodle 2.9.0
	CompressorLZHLW   Compressor = 1  // no longer supported as of Oodle 2.9.0
	CompressorLZNIB   Compressor = 2  // no longer supported as of Oodle 2.9.0
	CompressorLZBLW   Compressor = 5  // no longer supported as of Oodle 2.9.0
	CompressorLZA     Compressor = 6  // no longer supported as of Oodle 2.9.0
	*/

	CompressorCount   Compressor = 14
	CompressorForce32 Compressor = 0x40000000
)

// https://github.com/EpicGames/UnrealEngine/blob/46544fa5e0aa9e6740c19b44b0628b72e7bbd5ce/Engine/Source/Programs/Horde/HordeStorage/Jupiter.Common/Utils/OodleCompressor.cs#L37
type CompressionLevel int32

const (
	CompressionLevelNone      CompressionLevel = 0 // don't compress  just copy raw bytes
	CompressionLevelSuperFast CompressionLevel = 1 // super fast mode  lower compression ratio
	CompressionLevelVeryFast  CompressionLevel = 2 // fastest LZ mode with still decent compression ratio
	CompressionLevelFast      CompressionLevel = 3 // fast - good for daily use
	CompressionLevelNormal    CompressionLevel = 4 // standard medium speed LZ mode

	CompressionLevelOptimal1 CompressionLevel = 5 // optimal parse level 1 (faster optimal encoder)
	CompressionLevelOptimal2 CompressionLevel = 6 // optimal parse level 2 (recommended baseline optimal encoder)
	CompressionLevelOptimal3 CompressionLevel = 7 // optimal parse level 3 (slower optimal encoder)
	CompressionLevelOptimal4 CompressionLevel = 8 // optimal parse level 4 (very slow optimal encoder)
	CompressionLevelOptimal5 CompressionLevel = 9 // optimal parse level 5 (don't care about encode speed  maximum compression)

	CompressionLevelHyperFast1 CompressionLevel = -1 // faster than SuperFast  less compression
	CompressionLevelHyperFast2 CompressionLevel = -2 // faster than HyperFast1  less compression
	CompressionLevelHyperFast3 CompressionLevel = -3 // faster than HyperFast2  less compression
	CompressionLevelHyperFast4 CompressionLevel = -4 // fastest  less compression

	// aliases :

	CompressionLevelHyperFast CompressionLevel = CompressionLevelHyperFast1 // alias HyperFast base level
	CompressionLevelOptimal   CompressionLevel = CompressionLevelOptimal2   // alias optimal standard level
	CompressionLevelMax       CompressionLevel = CompressionLevelOptimal5   // maximum compression level
	CompressionLevelMin       CompressionLevel = CompressionLevelHyperFast4 // fastest compression level

	CompressionLevelForce32 CompressionLevel = 0x40000000
	CompressionLevelInvalid CompressionLevel = CompressionLevelForce32
)

// https://github.com/EpicGames/UnrealEngine/blob/46544fa5e0aa9e6740c19b44b0628b72e7bbd5ce/Engine/Source/Programs/Horde/HordeStorage/Jupiter.Common/Utils/OodleCompressor.cs#L66
type DecodeThreadPhase int32

const (
	DecodeThreadPhase1   DecodeThreadPhase = 1
	DecodeThreadPhase2   DecodeThreadPhase = 2
	DecodeThreadPhaseAll DecodeThreadPhase = 3
	DecodeUnthreaded     DecodeThreadPhase = DecodeThreadPhaseAll
)

// https://github.com/EpicGames/UnrealEngine/blob/46544fa5e0aa9e6740c19b44b0628b72e7bbd5ce/Engine/Source/Programs/Horde/HordeStorage/Jupiter.Common/Utils/OodleCompressor.cs#L74
type FuzzSafe int32

const (
	FuzzSafeNo  FuzzSafe = 0
	FuzzSafeYes FuzzSafe = 1
)

// https://github.com/EpicGames/UnrealEngine/blob/46544fa5e0aa9e6740c19b44b0628b72e7bbd5ce/Engine/Source/Programs/Horde/HordeStorage/Jupiter.Common/Utils/OodleCompressor.cs#L80
type CheckCRC int32

const (
	CheckCRCNo  CheckCRC = 0
	CheckCRCYes CheckCRC = 1

	CheckCRCForce32 CheckCRC = 0x40000000
)

// https://github.com/EpicGames/UnrealEngine/blob/46544fa5e0aa9e6740c19b44b0628b72e7bbd5ce/Engine/Source/Programs/Horde/HordeStorage/Jupiter.Common/Utils/OodleCompressor.cs#L87
type Verbosity int32

const (
	VerbosityNone    Verbosity = 0
	VerbosityMinimal Verbosity = 1
	VerbositySome    Verbosity = 2
	VerbosityLots    Verbosity = 3

	VerbosityForce32 Verbosity = 0x40000000
)
