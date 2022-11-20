package misc

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type EGuidFormats uint8

const (
	EGuidFormatDigits                         EGuidFormats = iota // 32 digits => eg: "00000000000000000000000000000000"
	EGuidFormatDigitsWithHyphens                                  // 32 digits separated by hyphens => eg: 00000000-0000-0000-0000-000000000000
	EGuidFormatDigitsWithHyphensInBraces                          // 32 digits separated by hyphens and enclosed in braces => eg: {00000000-0000-0000-0000-000000000000}
	EGuidFormatDigitsWithHyphensInParentheses                     // 32 digits separated by hyphens and enclosed in parentheses => eg: (00000000-0000-0000-0000-000000000000)
	EGuidFormatHexValuesInBraces                                  // Comma-separated hexadecimal values enclosed in braces => eg: {0x00000000,0x0000,0x0000,{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}}
	EGuidFormatUniqueObjectGuid                                   // This format is currently used by the FUniqueObjectGuid class => eg: 00000000-00000000-00000000-00000000
	EGuidFormatShort                                              // Base64 characters with dashes and underscores instead of pluses and slashes (respectively)=> eg: AQsMCQ0PAAUKCgQEBAgADQ
	EGuidFormatBase36Encoded                                      // Base-36 encoded, compatible with case-insensitive OS file systems (such as Windows) => eg: 1DPF6ARFCM4XH5RMWPU8TGR0J
)

var MainGUID = FGuid{}

// ue4\objects\core\misc\Guid.kt
type FGuid struct {
	A uint32 // Holds the first component
	B uint32 // Holds the second component
	C uint32 // Holds the third component
	D uint32 // Holds the fourth component
}

func NewFGuid(id string) FGuid {
	data, err := hex.DecodeString(id)
	if err != nil {
		panic(err)
	}

	var guid FGuid
	err = binary.Read(bytes.NewReader(data), binary.BigEndian, &guid)
	if err != nil {
		panic(err)
	}

	return guid
}

func (guid FGuid) IsValid() bool {
	return (guid.A | guid.B | guid.C | guid.D) != 0
}

func (guid FGuid) String(pattern EGuidFormats) string {
	switch pattern {
	case EGuidFormatDigitsWithHyphens:
		return fmt.Sprintf("%08X-%04X-%04X-%04X-%04X%08X", guid.A, guid.B>>16, guid.B&0xFFFF, guid.C>>16, guid.C&0xFFFF, guid.D)
	case EGuidFormatDigitsWithHyphensInBraces:
		return fmt.Sprintf("{%08X-%04X-%04X-%04X-%04X%08X}", guid.A, guid.B>>16, guid.B&0xFFFF, guid.C>>16, guid.C&0xFFFF, guid.D)
	case EGuidFormatDigitsWithHyphensInParentheses:
		return fmt.Sprintf("(%08X-%04X-%04X-%04X-%04X%08X)", guid.A, guid.B>>16, guid.B&0xFFFF, guid.C>>16, guid.C&0xFFFF, guid.D)
	case EGuidFormatHexValuesInBraces:
		return fmt.Sprintf(
			"{0x%08X,0x%04X,0x%04X,{0x%02X,0x%02X,0x%02X,0x%02X,0x%02X,0x%02X,0x%02X,0x%02X}}",
			guid.A, guid.B>>16, guid.B&0xFFFF,
			guid.C>>24, (guid.C>>16)&0xFF, (guid.C>>8)&0xFF, guid.C&0xFF,
			guid.D>>24, (guid.D>>16)&0xFF, (guid.D>>8)&0xFF, guid.D&0xFF,
		)
	case EGuidFormatUniqueObjectGuid:
		return fmt.Sprintf("%08X-%08X-%08X-%08X", guid.A, guid.B, guid.C, guid.D)
	}

	return fmt.Sprintf("%08x%08x%08x%08x", guid.A, guid.B, guid.C, guid.D)
}
