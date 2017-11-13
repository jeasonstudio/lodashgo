package Test

import (
	"lodashgo/String"
	"testing"
)

func Test_String_EndsWith(t *testing.T) {
	r1 := String.EndsWith("aabbcc", "bcc")
	r2 := String.EndsWith("aabbcc", "bc")
	r3 := String.EndsWith("aabbcc", "c")
	if r1 != true || r2 != false || r3 != true {
		t.Error("Error @ EndsWith")
	}
}
