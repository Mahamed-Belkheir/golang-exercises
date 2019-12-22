package exercises

import "fmt"

//Josephus ...
func Josephus(items []interface{}, k int) []interface{} {
	k--
	remaining := items
	remainingCount := len(remaining)
	dead := make([]interface{}, len(items))
	posOnGrave := 0
	turn := 0

	for remainingCount > 0 {
		turn = (turn + k) % len(remaining)
		dead[posOnGrave] = remaining[turn]
		posOnGrave++
		remaining = append(remaining[:turn], remaining[turn+1:]...)
		fmt.Println(remaining, dead)
		remainingCount--
	}

	return dead
}
