package main

type RESPPrefixBytes []byte

const (
	RESPPrefixBytesSimpleStrings = "+"
	RESPPrefixBytesError         = "-"
	RESPPrefixBytesIntegers      = ":"
	RESPPrefixBytesBulkStrings   = "$"
	RESPPrefixBytesArray         = "*"
)
