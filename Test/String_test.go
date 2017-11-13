package Test

import (
	"lodashgo/Util"
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

func Test_String_PadStart(t *testing.T) {
	r1 := String.PadStart("a", 3, "b")
	if r1 != "bba" { t.Error("Error @ PadStart") }
	r2 := String.PadStart("asdf", 2, "bc")
	if r2 != "asdf" { t.Error("Error @ PadStart") }
}

func Test_String_PadEnd(t *testing.T) {
	r1 := String.PadEnd("a", 3, "b")
	if r1 != "abb" { t.Error("Error @ PadEnd") }
	r2 := String.PadEnd("asdf", 2, "bc")
	if r2 != "asdf" { t.Error("Error @ PadEnd") }
}

func Test_String_UniqueID(t *testing.T) {
	d := Util.UniqueID("")
	a := Util.UniqueID("a")
	b := Util.UniqueID("b")
	if d() != "LODASHGO_DEFAULT_1" {
		t.Error("Error @ UniqueID")
	}
	if a() != "a1" {
		t.Error("Error @ UniqueID")
	}
	if a() != "a2" {
		t.Error("Error @ UniqueID")
	}
	if b() != "b1" {
		t.Error("Error @ UniqueID")
	}
	if a() != "a3" {
		t.Error("Error @ UniqueID")
	}
}
