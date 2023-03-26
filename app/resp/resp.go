package resp

import (
	"strings"
)

type prefix string

const (
	prefixSimpleStrings = "+"
	prefixError         = "-"
	prefixIntegers      = ":"
	prefixBulkStrings   = "$"
	prefixArray         = "*"
)

var delimiter = "\r\n"

func SimpleStrings(str string) []byte {
	return []byte(strings.Join([]string{prefixSimpleStrings, str}, delimiter) + delimiter)
}
