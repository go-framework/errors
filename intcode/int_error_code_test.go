package intcode

import (
	"testing"
)

func TestErrCode_WithDetail(t *testing.T) {
	t.Log(Succeed.WithDetail("Succeed"))
	t.Log(ErrFormat.WithDetail("ErrFormat"))
}

func TestGetCodeText(t *testing.T) {
	t.Log(GetCodeText(ErrUnmarshal))
	t.Log(GetCodeText(100))
}
