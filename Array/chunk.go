package Array

// Chunk Creates an array of elements split into groups the length of size. If array can't be split evenly, the final chunk will be the remaining elements.
func Chunk(s []interface{}, size int) ([][]interface{}, error) {
	result := [][]interface{} {}
	for index := 0; index < len(s); index += size {
		temp := [] interface{} {}
		for indexForSize := 0; indexForSize < size; indexForSize++ {
			if (index + indexForSize < len(s)) {
				temp = append(temp, s[index + indexForSize])
			}
		}
		result = append(result, temp)
	}
	return result, nil
}
