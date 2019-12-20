package exercises

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

//Arithmetic aaa
func Arithmetic(a int, b int, operator string) int {
	ops := map[string]interface{}{
		"add":      add,
		"subtract": sub,
		"multiply": mul,
		"divide":   div,
	}

	return ops[operator].(func(int, int) int)(a, b)
}
