package main

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (s *Circle) getType() string {
	return "Circle"
}
