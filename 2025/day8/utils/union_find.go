package utils 

type DSU struct {
	parent []int 
	size []int 
	Components int 
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size: make([]int, n),
		Components: n,
	}

	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	
	return dsu 
}

func (dsu *DSU) Find(v int) int {
	if v == dsu.parent[v] {
		return v 
	}
	dsu.parent[v] = dsu.Find(dsu.parent[v])
	return dsu.parent[v]
}

func (dsu *DSU) Union(v, u int) {
	rv, ru := dsu.Find(v), dsu.Find(u)

	if rv == ru {
		return 
	}

	if dsu.size[rv] < dsu.size[ru] {
		rv, ru = ru, rv 
	}
	dsu.Components--
	dsu.parent[ru] = rv 
	dsu.size[rv] += dsu.size[ru]
}