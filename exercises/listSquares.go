package exercises

import (
	"math"
)

//ListSquared ...
func ListSquared(m, n int) [][]int {
	result := [][]int{}
	for i := m; i < n; i++ {
		list := getArraySquared(i)
		sum := sum(list)
		if checkSquared(sum) {
			result = append(result, [][]int{[]int{i, sum}}...)
		}
	}

	return result
}

func checkSquared(a int) bool {
	root := math.Sqrt(float64(a))
	return root == float64(int64(root))
}

func sum(a []int) int {
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}

func getArraySquared(a int) []int {
	arr := []int{}
	for i := a; i > 0; i-- {
		if a%i == 0 {
			arr = append(arr, i*i)
		}
	}
	return arr
}
