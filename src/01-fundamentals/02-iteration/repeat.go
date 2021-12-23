package iteration

func Repeat(str string, n int) string {
	out := ""
	for i := 0; i < n; i += 1 {
		out += str
	}
	return out
}
