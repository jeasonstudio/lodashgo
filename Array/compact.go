package Array

// Compact Creates an array with all falsey values removed. The values false, nil and 0 are falsey.
func Compact(s []interface{}) ([]interface{}, error) {
	result := []interface{} {}
	for _, sItem := range s {
		if (sItem != nil && sItem != 0 && sItem != "" && sItem != false) {
			result = append(result, sItem)
		}
	}
	return result, nil
}
