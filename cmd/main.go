package main

import "fmt"

func main() {
	// This is a Batch Processing Example of Simple Pipeline.
	// ints := []int{1, 2, 3, 4}

	// multiply := func(values []int, multiplier int) []int {
	// 	multipliedValues := make([]int, len(values))
	// 	for i, v := range values {
	// 		multipliedValues[i] = v * multiplier
	// 	}
	// 	return multipliedValues
	// }

	// add := func(values []int, additive int) []int {
	// 	addedValues := make([]int, len(values))
	// 	for i, v := range values {
	// 		addedValues[i] = v + additive
	// 	}
	// 	return addedValues
	// }

	// for _, v := range multiply(add(ints, 1), 2) {
	// 	fmt.Println(v)
	// }

	multiply := func(val, multiplier int) int {
		return val * multiplier
	}

	add := func(val, additive int) int {
		return val + additive
	}

	ints := []int{1, 2, 3, 4}

	for _, v := range ints {
		fmt.Println(add(multiply(v, 2), 1))
	}

}
