package main

import (
	"fmt"
	"os"
	"strings"
	"day7/utils"
)

var rows, cols int
func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := strings.Split(strings.TrimSpace(string(f)), "\n")
	// there is \r at end of each row, remove it
	for i := range grid {
		grid[i] = strings.TrimRight(grid[i], "\r")
	}
	rows, cols = len(grid), len(grid[0])


	// part1(grid)
	part2(grid)
}

func part1(grid []string) {
	set := utils.NewSet[utils.Pair[int, int]]() // visited array
	que := utils.NewQueue[utils.Pair[int, int]]()
	split := utils.NewSet[utils.Pair[int, int]]() // split points

	for col := 0; col < cols; col++ {
		if grid[0][col] == 'S' {
			start := utils.Pair[int, int]{X: 0, Y: col}
			set.Add(start)
			que.Push(start)
			break
		}
	}
	
	for !que.Empty() {
		// breadth first search
		// traversing each level at a time
		size := que.Size()
		for i := 0; i < size; i++ {
			p := que.Pop()
			nx, ny := p.X + 1, p.Y // down
			if nx == rows {
				continue
			}

			down := utils.Pair[int, int]{X: nx, Y: ny}

			if grid[nx][ny] == '^' {
				split.Add(down) // split occured

				// LEFT
				if ny-1 >= 0 {
					left := utils.Pair[int, int]{X: nx, Y: ny - 1}
					if !set.Has(left) {
						set.Add(left)
						que.Push(left)
					}
				}

				// RIGHT
				if ny+1 < cols {
					right := utils.Pair[int, int]{X: nx, Y: ny + 1}
					if !set.Has(right) {
						set.Add(right)
						que.Push(right)
					}
				}
			} else {	
				if !set.Has(down) {
					set.Add(down)
					que.Push(down)
				}
			}
		}
	}
	fmt.Println("Part1: ", split.Size())
}


func part2(grid []string) {
	var start utils.Pair[int, int]
	for col := 0; col < cols; col++ {
		if grid[0][col] == 'S' {
			start = utils.Pair[int, int]{X: 0, Y: col}
		}
	}

	valid := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < rows && y < cols 
	}

	// recursion which will go L/R till it reaches last row
	// you must define function outside to use as recursion inside
	var f func(int, int) int
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	f = func(x, y int) int {
		if !valid(x, y) {
			return 0
		}
		if dp[x][y] != -1 {
			return dp[x][y]
		}
		// down
		nx, ny := x + 1, y 
		if nx == rows {
			return 1
		}

		if grid[nx][ny] == '^' {
			// split L + R
			dp[x][y] = f(nx, ny - 1) + f(nx, ny + 1)
			return dp[x][y]
		} 
		dp[x][y] = f(nx, ny)
		return f(nx, ny)
	}

	fmt.Println("Part2: ", f(start.X, start.Y))
}