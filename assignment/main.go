package main

import "fmt"

func main() {
	var intSlice []int
	for i := 0; i < 11; i++ {
		intSlice = append(intSlice, i)
	}
	for _, value := range intSlice {
		if value%2 == 0 {
			fmt.Printf("%d is even.\n", value)
		} else {
			fmt.Printf("%d is odd.\n", value)
		}
	}
}
