package Array

import (
	"errors"
)

// IntChunk Creates an array of elements split into groups the length of size.
// If array can't be split evenly, the final chunk will be the remaining elements.
func IntChunk(s []int, size int) ([][]int, error) {
	var result [][]int
	for index := 0; index < len(s); index += size {
		var temp []int
		for indexForSize := 0; indexForSize < size; indexForSize++ {
			if index+indexForSize < len(s) {
				temp = append(temp, s[index+indexForSize])
			}
		}
		result = append(result, temp)
	}
	return result, nil
}

// UnSafeChunk Creates an array of elements split into groups the length of size.
// If array can't be split evenly, the final chunk will be the remaining elements.
// The difference between Chunk is UnSafeChunk accept any kind type of Array.
func UnSafeChunk(s interface{}, size int) ([]interface{}, error) {
	var result []interface{}
	var err error
	// Types support
	switch s.(type) {
	case []int:
		selfSlice := s.([]int)
		for index := 0; index < len(selfSlice); index += size {
			temp := []int{}
			for indexForSize := 0; indexForSize < size; indexForSize++ {
				if index+indexForSize < len(selfSlice) {
					temp = append(temp, selfSlice[index+indexForSize])
				}
			}
			result = append(result, temp)
		}
	case []string:
		selfSlice := s.([]string)
		for index := 0; index < len(selfSlice); index += size {
			temp := []string{}
			for indexForSize := 0; indexForSize < size; indexForSize++ {
				if index+indexForSize < len(selfSlice) {
					temp = append(temp, selfSlice[index+indexForSize])
				}
			}
			result = append(result, temp)
		}
	case []interface{}:
		selfSlice := s.([]interface{})
		for index := 0; index < len(selfSlice); index += size {
			temp := []interface{}{}
			for indexForSize := 0; indexForSize < size; indexForSize++ {
				if index+indexForSize < len(selfSlice) {
					temp = append(temp, selfSlice[index+indexForSize])
				}
			}
			result = append(result, temp)
		}
	default:
		// Not match any kind of [int, string, interface{}]
		result = nil
		err = errors.New("Error: The first arg should be some kind of `Array`")
	}
	return result, err
}

// // 如果我想给 slice 写一个通用的方法函数 (比如 chunk)
// // 在没有泛型的情况下是不是只能想下面这样 blabla 每种类型都实现一遍?
// func IntChunk(s []int, size int) ([][]int, error)
// func Int8Chunk(s []int8, size int) ([][]int8, error)
// func Int32Chunk(s []int32, size int) ([][]int32, error)
// func Int64Chunk(s []int64, size int) ([][]int64, error)
// func UintChunk(s []uint, size int) ([][]uint, error)
// func Uint8Chunk(s []uint8, size int) ([][]uint8, error)
// func Uint32Chunk(s []uint32, size int) ([][]uint32, error)
// func Uint64Chunk(s []uint64, size int) ([][]uint64, error)
// func Float32Chunk(s []float32, size int) ([][]float32, error)
// func Float64Chunk(s []float64, size int) ([][]float64, error)
// func StringChunk(s []string, size int) ([][]string, error)

// // 实现成 interface{} 就丢失了静态类型检查, 感觉不太好 -_-
// func Chunk(s interface{}, size int) ([]interface{}, error)

// // 这样说对吗?
