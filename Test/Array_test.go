package Test

import (
	"lodashgo/Array"
	"testing"
)

func Test_Array_IntChunk(t *testing.T)  {
	t1 := []int {1, 2, 3, 4, 5, 7, 9}
	r1, _ := Array.IntChunk(t1, 2)
	r2, _ := Array.IntChunk(t1, 3)
	r3, _ := Array.IntChunk(t1, 4)
	r4, _ := Array.IntChunk(t1, 7)
	if len(r1) != 4 {
		t.Error("Some thing error")
	} else if len(r2) != 3 {
		t.Error("Some thing error")
	} else if len(r3) != 2 {
		t.Error("Some thing error")
	} else if len(r4) != 1 {
		t.Error("Some thing error")
	} else {
		t.Log("Test_Array_Chunk Passed")
	}
}

// func Test_Array_Compact(t *testing.T)  {
// 	t1 := []interface{} {1, false, nil, "sss", 0, true, ""}
// 	r, _ := Array.Compact(t1)
// 	if len(r) != 3 {
// 		t.Error("error")
// 	} else {
// 		t.Log("Test_Array_Compact Passed")
// 	}
// }
