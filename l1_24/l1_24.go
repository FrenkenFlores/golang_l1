package l1_24

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-q.X), 2) + math.Pow(float64(p.Y-q.Y), 2))
}
