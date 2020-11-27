package iteration

func Repeat(s string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += s
	}
	return repeated
}
