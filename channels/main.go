package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com"}

	c := make(chan string)
	for _, website := range websites {
		go checkLink(website, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			// if you use l directly the child routine and main routine
			// will be pointing to the variable and before after the sleep
			// command the value of l might be different.
			checkLink(link, c)
		}(l)
	}

}

func checkLink(website string, c chan string) {
	_, error := http.Get(website)
	if error != nil {
		fmt.Println("Getting error ", error, " when fetching ")
		c <- website
		return
	} else {
		fmt.Println("No errors for ", website)
	}
	c <- website
}
