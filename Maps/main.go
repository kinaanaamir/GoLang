package main

import "fmt"

func print(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println(color, hex)
	}
}
func main() {
	// var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"while": "#000000",
		"black": "#ffffff",
		"green": "#dfsdfe",
	}
	print(colors)
	//colors["red"] = "#ff0000"
	//delete(colors, "red")
	//fmt.Println(colors)
}
