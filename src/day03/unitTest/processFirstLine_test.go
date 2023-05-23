package unitTest

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestProcessFirstLine(t *testing.T) {
	firstLine := ProcessFirstLine()
	assert.Equal(t, "line00", firstLine)
}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}

func ReadFirstLine() string {
	open, err := os.Open("log")
	defer open.Close()
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
