package main

import "fmt"

type shape interface {
	getArea(height float64, base float64) float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	tr := triangle{}
	sq := square{}
	printArea(tr, 3.0, 10.0)
	printArea(sq, 33.0, 33.0)
}

func printArea(s shape, height float64, base float64 ) {
	fmt.Println(s.getArea(height, base))
}

func (triangle) getArea(height float64, base float64) float64 {
	return height * (base / 2)
}
func (square) getArea(sideLength float64, base float64) float64 {
	return sideLength * base
}