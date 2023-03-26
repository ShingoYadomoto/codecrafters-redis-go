package resp

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
	return []byte(prefixSimpleStrings + str + delimiter)
}
