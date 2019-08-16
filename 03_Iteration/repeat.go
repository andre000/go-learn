package iteration

const NUM_REPETITION = 5

func Repeat(char string) string {
	var repeated string
	for i := 0; i < NUM_REPETITION; i++ {
		repeated += char
	}
	return repeated
}
