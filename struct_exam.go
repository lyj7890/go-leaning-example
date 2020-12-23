package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p *Point) show() {
	fmt.Println(p.x, p.y)
}

type Circle struct {
	*Point
	Radius int
}

func main() {
	var p = &Point{x: 10, y: 20}
	c := Circle{p, 30}
	c.show()
	c.Point.show()
}
