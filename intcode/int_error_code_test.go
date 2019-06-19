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

func TestErrCode_Level(t *testing.T) {
	t.Log(ErrNil.Normal("Normal"))
	t.Log(ErrNil.Debug("Debug"))
	t.Log(ErrNil.Warn("Warn"))
	t.Log(ErrNil.Critical("Critical"))

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Log("get recover")
				t.Log("panic", r)
			}
		}()

		t.Log(ErrNil.Panic("Panic"))
	}()

}

func TestErrCode_Level_Fatal(t *testing.T) {
	t.Log(ErrNil.Fatal("Fatal"))
}
