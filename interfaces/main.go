package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishbot struct{}
type spanishbot struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishbot) getGreeting() string {
	return "Hi there"
}

func (spanishbot) getGreeting() string {
	return "Hola"
}
