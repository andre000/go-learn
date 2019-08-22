package iteration

// Repeat prints the given word {loop} times
func Repeat(char string, loop int) string {
	var repeated string
	for i := 0; i < loop; i++ {
		repeated += char
	}
	return repeated
}
