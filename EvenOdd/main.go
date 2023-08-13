package main

import "fmt"

func main() {
	array := []int{}
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	for _, value := range array {
		if value%2 == 0 {
			fmt.Printf("%v is Even\n", value)
		} else {
			fmt.Printf("%v is Odd\n", value)
		}
	}
}
