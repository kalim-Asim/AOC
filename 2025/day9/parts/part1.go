package parts

import (
	"fmt"
	"math"
	. "day9/utils"
)

func Part1(points []Point) {
	n, maxArea := len(points), 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			area := Area(points[i], points[j])
			maxArea = max(maxArea, area)
		}
	}

	fmt.Println("Part1: ", maxArea)
}

func Area(a, b Point) int {
	dx := int(math.Abs(float64(a.X-b.X)))
	dy := int(math.Abs(float64(a.Y - b.Y)))
	return (dx + 1) * (dy + 1)
}