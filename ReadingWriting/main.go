package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) <= 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	filename := os.Args[1]
	file, error := os.Open(filename)

	if error != nil {
		fmt.Println("Following error occurred", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
