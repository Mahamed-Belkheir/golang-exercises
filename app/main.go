package main

import (
	"fmt"
	"golang-exercises/exercises"
)

func main() {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(exercises.Comp(a1, a2))
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144}
	fmt.Println(exercises.Comp(a1, a2))
}
