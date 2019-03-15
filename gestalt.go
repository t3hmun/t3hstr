package t3hstr

import "fmt"

// Gestalt truncates a line to maxlen by removing the middle instead of the end.
// Puts separator in the middle when runes are removed (without overflowing maxlen).
// Undefined result behaviour when newlines are present.
func Gestalt(target []rune, maxlen int, separator []rune) string {
	tlen := len(target)
	if tlen <= maxlen {
		return string(target)
	}

	seplen := len(separator)
	contentlen := (maxlen - seplen) / 2
	stopBefore := contentlen
	startAt := tlen - (contentlen)
	gestalt := fmt.Sprintf("%s%s%s", string(target[:stopBefore]), string(separator), string(target[startAt:]))
	return gestalt
}
