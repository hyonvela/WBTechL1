package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x float64, y float64) Point {
	return Point{x, y}
}

func (p Point) calculateDistance(p2 Point) float64 {
	return math.Sqrt((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y))
}

func main() {
	p1 := newPoint(-6.6, 3.5)
	p2 := newPoint(3.75, 4.1)

	fmt.Println(p1.calculateDistance(p2))
}
