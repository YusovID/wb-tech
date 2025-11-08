package point

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(otherPoint Point) float64 {
	// d = √((x2 - x1)² + (y2 - y1)²)
	dx := otherPoint.x - p.x
	dy := otherPoint.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}
