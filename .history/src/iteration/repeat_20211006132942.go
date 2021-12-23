package iteration

func Repeat(str string) string {
	out := ""
	for i := 0; i < 5; i += 1 {
		out += str
	}
	return out
}
