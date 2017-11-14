package Util

import "testing"

func Test_String_UniqueID(t *testing.T) {
	d := UniqueID("")
	a := UniqueID("a")
	b := UniqueID("b")
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
