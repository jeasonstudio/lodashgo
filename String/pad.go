package String

import (
	"strings"
)

// PadStart Pads string on the left side if it's shorter than length.
// Padding characters are truncated if they exceed length.
func PadStart(source string, length int, chars string) string {
	sourceSlice := strings.Split(source, "")
	// no need to pad
	if length <= len(sourceSlice) { return source }

	return getLengthString(chars, length - len(sourceSlice)) + source
}

// PadEnd Pads string on the right side if it's shorter than length.
// Padding characters are truncated if they exceed length.
func PadEnd(source string, length int, chars string) string {
	sourceSlice := strings.Split(source, "")
	// no need to pad
	if length <= len(sourceSlice) { return source }

	return source + getLengthString(chars, length - len(sourceSlice))
}

func getLengthString(s string, l int) string {
	charsSlice := strings.Split(s, "")
	var padString []string
	for index := 0; index < l; index++ {
		padString = append(padString, charsSlice[index % len(charsSlice)])
	}
	return strings.Join(padString, "")
}