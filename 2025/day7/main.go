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


	part1(grid)
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