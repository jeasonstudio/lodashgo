package Array

// IntReduce collection to a value which is the accumulated result of running each element in collection thru iteratee,
// where each successive invocation is supplied the return value of the previous.
// If accumulator is not given, the first element of collection is used as the initial value.
// The iteratee is invoked with four arguments:
func IntReduce(source []int, f func(accumulator interface{}, currentValue, currentIndex int, array []int) interface{}) (interface{}, error) {
	var accumulator interface{} = source[0]
	for index := 1; index < len(source); index++ {
		accumulator = f(accumulator, source[index], index, source)
	}
	return accumulator, nil
}

// StringReduce collection to a value which is the accumulated result of running each element in collection thru iteratee,
// where each successive invocation is supplied the return value of the previous.
// If accumulator is not given, the first element of collection is used as the initial value.
// The iteratee is invoked with four arguments:
func StringReduce(source []string, f func(accumulator interface{}, currentValue string, currentIndex int, array []string) interface{}) (interface{}, error) {
	var accumulator interface{} = source[0]
	for index := 1; index < len(source); index++ {
		accumulator = f(accumulator, source[index], index, source)
	}
	return accumulator, nil
}
