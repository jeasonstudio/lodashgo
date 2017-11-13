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
			if (index + indexForSize < len(s)) {
				temp = append(temp, s[index + indexForSize])
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
					temp := []int {}
					for indexForSize := 0; indexForSize < size; indexForSize++ {
						if (index + indexForSize < len(selfSlice)) {
							temp = append(temp, selfSlice[index + indexForSize])
						}
					}
					result = append(result, temp)
				}
			case []string:
				selfSlice := s.([]string)
				for index := 0; index < len(selfSlice); index += size {
					temp := []string {}
					for indexForSize := 0; indexForSize < size; indexForSize++ {
						if (index + indexForSize < len(selfSlice)) {
							temp = append(temp, selfSlice[index + indexForSize])
						}
					}
					result = append(result, temp)
				}
			case []interface{}:
				selfSlice := s.([]interface{})
				for index := 0; index < len(selfSlice); index += size {
					temp := []interface{} {}
					for indexForSize := 0; indexForSize < size; indexForSize++ {
						if (index + indexForSize < len(selfSlice)) {
							temp = append(temp, selfSlice[index + indexForSize])
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
