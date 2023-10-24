package iteration

// Repeats the passed string N times.
func Repeat(character string, repeats int) string {
	var repeated string

	for i := 0; i < repeats; i++ {
		repeated += character
	}

	return repeated
}
