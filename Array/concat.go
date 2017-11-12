package Array

// Concat Creates a new array concatenating array with any additional arrays and/or values.
func Concat(o... interface{}) ([]interface{}, error) {
	result := [] interface {} {}
	for _, ot := range o {
		switch t := ot.(type) {
			case []interface{}:
				for _, val := range t {
					result = append(result, val)
				}
			case []string:
				for _, val := range t {
					result = append(result, val)
				}
			case []int:
				for _, val := range t {
					result = append(result, val)
				}
			default:
				result = append(result, ot)
		}
	}
	return result, nil
}
