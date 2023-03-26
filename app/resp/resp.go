package resp

const (
	prefixSimpleStrings = "+"
	prefixError         = "-"
	prefixIntegers      = ":"
	prefixBulkStrings   = "$"
	prefixArray         = "*"

	delimiter = "\r\n"
)

func SimpleStrings(str string) []byte {
	return []byte(prefixSimpleStrings + str + delimiter)
}
