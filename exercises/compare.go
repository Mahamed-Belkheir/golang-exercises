package exercises

//Comp ...
func Comp(array1 []int, array2 []int) bool {

	if array1 == nil || array2 == nil || len(array1) != len(array2) {
		return false
	}

	counter := 0
	mapOfArrSquared := make(map[int]int, len(array1))
	for _, num := range array1 {
		mapOfArrSquared[num*num]++
	}

	for _, num := range array2 {
		if mapOfArrSquared[num] > 0 {
			mapOfArrSquared[num]--
			counter++
		}
	}

	if counter == len(array1) {
		return true
	}

	return false

}
