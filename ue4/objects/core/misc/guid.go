package misc

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
