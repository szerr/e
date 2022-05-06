package e

import (
	"testing"
)

func TestAll(t *testing.T) {
	err := NewErr("err")
	o := NewErro("o", err)
	if IfErr(err) {
		t.Log("if ok")
	}
	if Trace(err) {
		t.Log("trace ok")
	}
	if AsInfo(o, err) {
		t.Log("as ok")
	}
	if IsInfos(o, err) {
		t.Log("ls ok")
	}
	defer func() {
		if err := recover(); err != nil {
			t.Log("fatal is ok")
		}
	}()
	if Fatals(err) {
		t.Fatalf("panic not recover")
	}
}
