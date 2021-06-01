package iteration

// Repeat function returns string
// contains character repeated n times
func Repeat(character string, n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += character
	}
	return res
}
