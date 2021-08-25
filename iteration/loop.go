package iteration

func Repeat(arg string) string {
	var repeated string

	for i := 0; i < 4; i++ {
		repeated += arg
	}

	return repeated
}