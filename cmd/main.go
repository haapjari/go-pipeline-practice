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

	// This is the Stream Processing Example

	// multiply := func(val, multiplier int) int {
	// 	return val * multiplier
	// }

	// add := func(val, additive int) int {
	// 	return val + additive
	// }

	// ints := []int{1, 2, 3, 4}

	// for _, v := range ints {
	// 	fmt.Println(add(multiply(v, 2), 1))
	// }

	// Stream, but with Channels Example
	generator := func(done <-chan any, integers ...int) <- chan int {
		intStream := make(chan int)

		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
				return
				case intStream <- i:
				}
			}
		}()

		return intStream
	}

	multiply := func(done <-chan any, intStream <-chan int, multiplier int) <- chan int{
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
				return
				case multipliedStream <- i*multiplier:
				}
			}
		}()
		return multipliedStream
	}

	done := make (chan any)
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, intStream, 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
