package Test

import (
	"lodashgo/Array"
	"testing"
)

func Test_Array_IntChunk(t *testing.T) {
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

func Test_Array_UnSafeChunk(t *testing.T) {
	t1 := []int {1, 2, 3, 4, 5, 7, 9}
	t2 := []string {"1", "2", "3", "4", "5", "7", "9"}
	t3 := []interface{} {1, 2, "jjj", nil, 5, 7, 9}
	t4 := []int32 {1, 2, 3, 4, 5, 7, 9}
	r1, e1 := Array.UnSafeChunk(t1, 2)
	if len(r1) != 4 || e1 != nil {
		t.Error("Error @ UnSafeChunk int array")
	}
	r2, e2 := Array.UnSafeChunk(t2, 2)
	if len(r2) != 4 || e2 != nil {
		t.Error("Error @ UnSafeChunk string array")
	}
	r3, e3 := Array.UnSafeChunk(t3, 2)
	if len(r3) != 4 || e3 != nil {
		t.Error("Error @ UnSafeChunk interface{} array")
	}
	_, e4 := Array.UnSafeChunk(t4, 2)
	if e4 == nil {
		t.Error("Error @ UnSafeChunk unsupported array")
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
