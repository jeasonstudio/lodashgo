package Array

// Concat Creates a new array concatenating array with any additional arrays and/or values.
func Concat(o... interface{}) ([]interface{}, error) {
	result := [] interface {} {}
	for _, ot := range o {
		switch t := ot.(type) {
			default:
				result = append(result, ot)
			case []interface{}:
				for _, val := range t {
					result = append(result, val)
				}
		}
	}
	return result, nil
}
