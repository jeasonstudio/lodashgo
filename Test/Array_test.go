package Test

import (
	"lodashgo/Array"
	"testing"
)

func Test_Array_Chunk(t *testing.T)  {
	t1 := []interface{} {1, 2, nil, "sss", true}
	r1, _ := Array.Chunk(t1, 2)
	r2, _ := Array.Chunk(t1, 3)
	r3, _ := Array.Chunk(t1, 4)
	r4, _ := Array.Chunk(t1, 7)
	if len(r1) != 3 {
		t.Error("Some thing error")
	} else if len(r2) != 2 {
		t.Error("Some thing error")
	} else if len(r3) != 2 {
		t.Error("Some thing error")
	} else if len(r4) != 1 {
		t.Error("Some thing error")
	} else {
		t.Log("Test_Array_Chunk Passed")
	}
}

func Test_Array_Compact(t *testing.T)  {
	t1 := []interface{} {1, false, nil, "sss", 0, true, ""}
	r, _ := Array.Compact(t1)
	if len(r) != 3 {
		t.Error("error")
	} else {
		t.Log("Test_Array_Compact Passed")
	}
}
