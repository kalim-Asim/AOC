package parts

import (
	"day11/utils"
	"fmt"
)

func Part2(edges []utils.Edge) {
	G, mp := utils.BuildGraph2(edges)

	src := mp["svr"]
	dst := mp["out"]
	fft := mp["fft"]
	dac := mp["dac"]

	ok := canReachDst(G, dst)

	// prune graph
	for u := range G {
		// these nodes can't reach dst
		// just remove them from our graph
		if !ok[u] {
			G[u] = nil
		}
	}

	comp, sccCount := tarjan(G)

	// build DAG
	DAG := make([][]int, sccCount)
	mask := make([]int, sccCount)

	if ok[fft] {
		mask[comp[fft]] |= 1
	}
	if ok[dac] {
		mask[comp[dac]] |= 2
	}

	for u := range G {
		if !ok[u] {
			continue
		}
		for _, v := range G[u] {
			if ok[v] && comp[u] != comp[v] {
				DAG[comp[u]] = append(DAG[comp[u]], comp[v])
			}
		}
	}

	srcC := comp[src]
	dstC := comp[dst]

	// DP on DAG
	dp := make([][4]int, sccCount)
	dp[srcC][mask[srcC]] = 1

	order := topo(DAG)

	for _, u := range order {
		for m := 0; m < 4; m++ {
			if dp[u][m] == 0 {
				continue
			}
			for _, v := range DAG[u] {
				nm := m | mask[v]
				dp[v][nm] += dp[u][m]
			}
		}
	}

	fmt.Println("Part2:", dp[dstC][3])
}

func reverseGraph(G [][]int) [][]int {
	n := len(G)
	R := make([][]int, n)
	for u := range G {
		for _, v := range G[u] {
			R[v] = append(R[v], u)
		}
	}
	return R
}

func canReachDst(G [][]int, dst int) []bool {
	R := reverseGraph(G)
	n := len(G)
	ok := make([]bool, n)

	stack := []int{dst}
	ok[dst] = true

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, v := range R[u] {
			if !ok[v] {
				ok[v] = true
				stack = append(stack, v)
			}
		}
	}
	return ok
}

func tarjan(G [][]int) ([]int, int) {
	n := len(G)
	disc := make([]int, n)
	low := make([]int, n)
	onStack := make([]bool, n)
	stack := []int{}
	for i := range disc {
		disc[i] = -1
	}

	time, sccID := 0, 0
	comp := make([]int, n)

	var dfs func(int)
	dfs = func(u int) {
		disc[u] = time
		low[u] = time
		time++
		stack = append(stack, u)
		onStack[u] = true

		for _, v := range G[u] {
			if disc[v] == -1 {
				dfs(v)
				low[u] = min(low[u], low[v])
			} else if onStack[v] {
				low[u] = min(low[u], disc[v])
			}
		}

		if low[u] == disc[u] {
			for {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[v] = false
				comp[v] = sccID
				if v == u {
					break
				}
			}
			sccID++
		}
	}

	for i := 0; i < n; i++ {
		if disc[i] == -1 {
			dfs(i)
		}
	}

	return comp, sccID
}

func topo(G [][]int) []int {
	n := len(G)
	ind := make([]int, n)
	for u := range G {
		for _, v := range G[u] {
			ind[v]++
		}
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if ind[i] == 0 {
			q = append(q, i)
		}
	}

	order := []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		order = append(order, u)
		for _, v := range G[u] {
			ind[v]--
			if ind[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return order
}

/* 		didn't worked
	// need to paas these nodes
// var fft, dac int 
// func Part2(graph []utils.Edge) {
// 	G, mp, counter := utils.BuildGraph(graph)
// 	fft, dac = mp["fft"], mp["dac"]
// 	src, dst := mp["svr"], mp["out"]
// 	ans := dfs2(src, dst, G, make([]bool, counter))
// 	fmt.Println("Part2: ", ans)
// }
// var calls int64

// func dfs2(cur, dst int, G [][]int, visited []bool) int {
// 	calls++
// 	if calls % 1_000_000 == 0 {
// 			fmt.Println("calls:", calls)
// 	}

// 	visited[cur] = true
// 	defer func() { visited[cur] = false }()

// 	if cur == dst {
// 		if visited[fft] && visited[dac] {
// 			return 1
// 		}
// 		return 0
// 	}
// 	total := 0

// 	for _, neigh := range G[cur] {
// 		if visited[neigh] {
// 			continue 
// 		}
// 		total += dfs2(neigh, dst, G, visited)
// 	}
// 	return total
// }

*/