package utils

var counter int
var mp map[string]int
// maps string nodes to numbers
func BuildGraph(graph []Edge) ([][]int, map[string]int, int) {
	mp = make(map[string]int)
	counter := 0
	
	for _, edge := range graph {
		from, to_s := edge.From, edge.To
		if _, ok := mp[from]; !ok {
			mp[from] = counter
			counter++
		}

		for _, to := range to_s {
			if _, ok := mp[to]; !ok {
				mp[to] = counter 
				counter++
			}

		}
	}

	G := make([][]int, counter)
	for _, edge := range graph {
		from, to_s := mp[edge.From], edge.To 

		for _, v := range to_s {
			to := mp[v]
			G[from] = append(G[from], to)
		}
	}
	return G, mp, counter
}