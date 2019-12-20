package exercises

import (
	"strconv"
	"strings"
)

//HighAndLow ...
func HighAndLow(in string) string {
	array := strings.Fields(in)
	lowest, _ := strconv.Atoi(array[0])
	highest, _ := strconv.Atoi(array[0])

	for i := 1; i < len(array); i++ {
		cur, _ := strconv.Atoi(array[i])
		if cur < lowest {
			lowest = cur
		}
		if cur > highest {
			highest = cur
		}
	}

	return strconv.Itoa(highest) + " " + strconv.Itoa(lowest)

}
