package util

import (
	"encoding/hex"
	"strings"
	"unicode"
)

// ConvertToHex ...
func ConvertToHex(s string) string {

	s = strings.TrimSpace(strings.ToLower(s))
	b := []byte(s)
	result := hex.EncodeToString(b)
	// t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	// result, _, err := transform.String(t, s)
	// if err != nil {
	// 	panic(err)
	// }
	// result = strings.ReplaceAll(result, "đ", "d")
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
