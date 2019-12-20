package exercises

import (
	"strconv"
	"strings"
)

//NbDig ...
func NbDig(n int, d int) int {
	// array := make([]string, n+1)
	counter := 0
	checker := strconv.Itoa(d)

	for i := 0; i <= n; i++ {
		entry := strconv.Itoa(i * i)
		entryArray := strings.Split(entry, "")
		for _, ch := range entryArray {
			if ch == checker {
				counter++
			}
		}
	}
	return counter
}
