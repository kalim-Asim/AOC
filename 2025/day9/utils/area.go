package utils

import "math"

func Area(a, b Point) int {
	dx := int(math.Abs(float64(a.X-b.X)))
	dy := int(math.Abs(float64(a.Y - b.Y)))
	return (dx + 1) * (dy + 1)
}