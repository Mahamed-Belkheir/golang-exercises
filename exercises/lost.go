package exercises

//Maps ...
func Maps(x []int) []int {
	b := make([]int, len(x))
	copy(b, x)
	for i := 0; i < len(x); i++ {
		b[i] = b[i] * 2
	}
	return b
}
