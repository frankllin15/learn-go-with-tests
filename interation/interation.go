package interation

// Repeat takes two arguments, a string and a int and return the string repeated the number of times of the second argument
func Repeat(character string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += character

	}
	return
}
