package iteration

// Basic looping function
func Repeat(arg string, n int) string {
	var repeated string

	for i := 0; i < n; i++ {
		repeated += arg
	}

	return repeated
}