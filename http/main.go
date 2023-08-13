package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(buffer []byte) (int, error) {
	fmt.Println(string(buffer))
	fmt.Println("Just wrote this many bytes", len(buffer))
	return len(buffer), nil
}
