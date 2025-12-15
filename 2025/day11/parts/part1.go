package parts

import (
	"day11/utils"
	"fmt"
)


func Part1(graph []utils.Edge) {
	G, mp, counter := utils.BuildGraph(graph)
	src, dst := mp["you"], mp["out"]
	ans := dfs(src, dst, G, make([]bool, counter))
	fmt.Println("Part1: ", ans)
}

// counts unique path from src to dest
func dfs(cur int, dst int, G [][]int, visited []bool) int {
	if cur == dst {
		return 1
	}
	visited[cur] = true 

	total := 0
	for _, neigh := range G[cur] {
		if visited[neigh] {
			continue 
		}
		total += dfs(neigh, dst, G, visited)
	}

	visited[cur] = false
	return total
}

