package unitTest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func JudgePassLine(score int16) bool {
	if score >= 60 {
		return true
	}
	return false
}

func TestJudgePassLineTrue(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}

func TestJudgePassLineFail(t *testing.T) {
	isPass := JudgePassLine(50)
	assert.Equal(t, false, isPass)
}
