package resp

import (
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

	delimiter = `\r\n`
)

func SimpleStrings(str string) []byte {
	return []byte(prefixSimpleStrings + str + delimiter)
}

const (
	commandPing = "PING"
	commandEcho = "ECHO"
)

type command struct {
	cmd  string
	args []string
}

func (c *command) Cmd() string {
	return c.cmd
}

// ParseCommand supports ECHO and PING commands only
func ParseCommand(b []byte) (*command, error) {
	strList := strings.Split(string(b), delimiter)
	strList = strList[:len(strList)-1] // 末尾は空文字になるので削除

	dataType, dataLenStr, strList := strList[0][:1], strList[0][1:], strList[1:]
	if dataType != prefixArray {
		return nil, fmt.Errorf("unexpected data type: %s", dataType)
	}

	dataLen, err := strconv.Atoi(dataLenStr)
	if err != nil {
		return nil, fmt.Errorf("unexpected data length. err: %s", err.Error())
	}

	expectLen, actualLen := 2*dataLen, len(strList)
	if actualLen != expectLen {
		return nil, fmt.Errorf("unexpected data length. expected: %d, actual: %d", expectLen, actualLen)
	}

	args := make([]string, dataLen)
	for i := 1; i < len(strList); i += 2 {
		args[i/2] = strList[i]
	}

	cmd := strings.ToUpper(args[0])
	validCmd := map[string]struct{}{
		commandPing: {},
		commandEcho: {},
	}

	if _, valid := validCmd[cmd]; !valid {
		return nil, fmt.Errorf("invalid command. %s", cmd)
	}

	return &command{
		cmd:  cmd,
		args: args[1:],
	}, nil
}
