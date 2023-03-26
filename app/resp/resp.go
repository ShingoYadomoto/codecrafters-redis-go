package resp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	prefixSimpleStrings = "+"
	prefixError         = "-"
	prefixIntegers      = ":"
	prefixBulkStrings   = "$"
	prefixArray         = "*"

	delimiter = "\r\n"

	commandPing = "PING"
	commandEcho = "ECHO"
)

func addEndDelimiter(str string) string {
	if strings.HasSuffix(str, delimiter) {
		return str
	}

	return str + delimiter
}

func simpleStrings(str string) []byte {
	return []byte(addEndDelimiter(prefixSimpleStrings + str))
}

func join(strl []string) []byte {
	return []byte(addEndDelimiter(strings.Join(strl, delimiter)))
}

func array(str string, len int) []byte {
	l := []string{
		prefixArray + fmt.Sprint(len),
		str,
	}
	return join(l)
}

var (
	ErrInvalidCommand = errors.New("invalid command")

	validCmd = map[string]struct{}{
		commandPing: {},
		commandEcho: {},
	}
)

type command struct {
	cmd     string
	argsStr string
	argsLen int
}

func (c *command) ping() ([]byte, error) {
	if c.argsStr == "" {
		return simpleStrings("PONG"), nil
	}
	return array(c.argsStr, c.argsLen), nil
}

func (c *command) echo() ([]byte, error) {
	return []byte(c.argsStr), nil
}

func (c *command) Response() ([]byte, error) {
	switch c.cmd {
	case commandPing:
		return c.ping()
	case commandEcho:
		return c.echo()
	}
	return nil, ErrInvalidCommand
}

// ParseCommand supports ECHO and PING commands only
func ParseCommand(b []byte) (*command, error) {
	var (
		str        = strings.TrimRight(string(b), "\x00")
		strList    = strings.Split(str, delimiter)
		dataHeader = strList[0]
		cmd        = strings.ToUpper(strList[2])
	)

	err := func() error {
		dataType, dataLenStr := dataHeader[:1], dataHeader[1:]
		if dataType != prefixArray {
			return fmt.Errorf("unexpected data type: %s", dataType)
		}

		dataLen, err := strconv.Atoi(dataLenStr)
		if err != nil {
			return fmt.Errorf("unexpected data length. err: %s", err.Error())
		}

		expectLen, actualLen := 2*dataLen+2, len(strList)
		if actualLen != expectLen {
			return fmt.Errorf("unexpected data length. expected: %d, actual: %d", expectLen, actualLen)
		}

		if _, valid := validCmd[cmd]; !valid {
			return fmt.Errorf("invalid command. %s", cmd)
		}

		return nil
	}()
	if err != nil {
		return nil, fmt.Errorf("rrong command usage. err: %s", err.Error())
	}

	var (
		argsStr = ""
		argsLen = 0
	)
	if len(strList) > 2 {
		var (
			argList  = strList[3:]
			startArg = argList[0]
		)

		if startArg != "" {
			argsStr = strings.Join(argList, delimiter)
			argsLen = len(argList) / 2
		}
	}

	return &command{
		cmd:     cmd,
		argsStr: argsStr,
		argsLen: argsLen,
	}, nil
}
