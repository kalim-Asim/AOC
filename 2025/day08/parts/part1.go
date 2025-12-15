package parts

import (
	"fmt"
	"sort"
	"day8/utils"
)

/*
	1. everytime find the two points which are shortest ?? optimize ?
	2. then check if belong to same group
	3. if same group, nothing
	4. if different union

	to do step 1. maintain an 2d array with [distance, i, j] sort in asc order O(N*N*logN)
	UNION FIND
*/

// do operation 1000 times
var COUNT int = 1000
func Part1(points []utils.Point) {
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
	for COUNT > 0 {
		_, v, u := arr[index][0], arr[index][1], arr[index][2]
		root_v, root_u := dsu.Find(v), dsu.Find(u)
		
		index++
		COUNT--
		if root_u == root_v {
			continue
		}

		dsu.Union(root_v, root_u)
	}

	// find 3 largest group ?
	comp := map[int]int{}
	for i := 0; i < n; i++ {
		root := dsu.Find(i)
		comp[root]++
	}

	sizes := make([]int, 0, len(comp))
	for _, sz := range comp {
		sizes = append(sizes, sz)
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	top3 := sizes[:3]
	var ans = top3[0] * top3[1] * top3[2]
	fmt.Println("Part1: ", ans)
}