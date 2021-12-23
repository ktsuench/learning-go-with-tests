package iteration

func Repeat(str string) string {
	i := 0
	out := ""
	for i < 5 {
		out += str
		i += 1
	}
	return out
}
