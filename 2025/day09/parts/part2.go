package parts

import (
	"fmt"
	. "day9/utils"
)

// mpx -> for same x what are possible y
// mpy -> for same y what are possible x
var mpx, mpy map[int][]int

func Part2(points []Point) {
	mpx = make(map[int][]int)
	mpy = make(map[int][]int)

	n, maxArea := len(points), 0

	for _, point := range points {
		x, y := point.X, point.Y 
		mpx[x] = append(mpx[x], y)
		mpy[y] = append(mpy[y], x)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !valid(points[i], points[j]) {
				continue
			}
			area := Area(points[i], points[j])
			maxArea = max(maxArea, area)
		}
	}
	fmt.Println("Part2: ", maxArea)
}

// checks if rectangle formed by p1 and p2 is valid
func valid(p1, p2 Point) bool {
	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y

	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	// top & bottom edges
	if !hasBothSides(mpy[y1], x1, x2) || !hasBothSides(mpy[y2], x1, x2) {
		return false
	}
	
	// left & right edges
	if !hasBothSides(mpx[x1], y1, y2) || !hasBothSides(mpx[x2], y1, y2){
		return false
	}
	
	return true
}


func hasBothSides(arr []int, low, high int) bool {
	left, right := false, false
	for _, v := range arr {
		if v <= low {
			left = true
		}
		if v >= high {
			right = true
		}
	}
	fmt.Println(left, right)
	return left && right
}
