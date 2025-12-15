package utils

func BuildGraph2(edges []Edge) ([][]int, map[string]int) {
	mp := make(map[string]int)
	id := 0

	getID := func(s string) int {
		if v, ok := mp[s]; ok {
			return v
		}
		mp[s] = id
		id++
		return mp[s]
	}

	for _, e := range edges {
		getID(e.From)
		for _, v := range e.To {
			getID(v)
		}
	}

	G := make([][]int, id)
	for _, e := range edges {
		u := mp[e.From]
		for _, v := range e.To {
			G[u] = append(G[u], mp[v])
		}
	}

	return G, mp
}
