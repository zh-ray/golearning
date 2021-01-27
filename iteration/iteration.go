package iteration

const repeatCount = 5

// Repeat takes a string and repeat several times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
