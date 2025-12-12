package parts

import (
	"day8/utils"
	"sort"
	"fmt"
)
/*
	unite every vertices till it gets into
	1 big component
	store last 2 vertices that u unite
*/
func Part2(points []utils.Point) {
	n := len(points)
	arr := make([][3]int, 0)
	dsu := utils.NewDSU(n)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := utils.Distance(points[i], points[j])
			arr = append(arr, [3]int{dist, i, j})
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	index := 0

	var ans int
	for dsu.Components != 1 {
		// fmt.Println(dsu.Components)
		_, v, u := arr[index][0], arr[index][1], arr[index][2]
		rv, ru := dsu.Find(v), dsu.Find(u) 
		index++

		if ru == rv {
			continue
		}

		dsu.Union(rv, ru)

		if dsu.Components == 1 {
			ans = points[v].X * points[u].X 
			break
		}
	}

	fmt.Println("Part2: ", ans, dsu.Components)
}
