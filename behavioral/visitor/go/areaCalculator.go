package main

import (
	"fmt"
	"math"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	a.area = s.side * s.side
	fmt.Println("Calculating area fir square")
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	a.area = int(math.Pi * float64(c.radius) * float64(c.radius))
	fmt.Println("Calculating area fir circle")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	a.area = r.b * r.l
	fmt.Println("Calculating area fir rectangle")
}
