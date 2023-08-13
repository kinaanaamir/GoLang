package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shapes interface {
	getArea() float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return tr.base * tr.height * 0.5
}

func printArea(s shapes) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{4}
	tr := triangle{4, 4}
	fmt.Println("square")
	printArea(sq)
	fmt.Println("\ntriangle")
	printArea(tr)
}
