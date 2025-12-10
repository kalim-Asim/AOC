package main

import (
	"fmt"
	"os"
	"strings"
)

var rows, cols int 
var dirs = [][]int{
	{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	{-1, 1}, {-1, -1}, {1, -1}, {1, 1},
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := strings.Split(strings.TrimSpace(string(file)), "\n")
	for i := range grid {
		grid[i] = strings.TrimRight(grid[i], "\r")
	}
	rows, cols = len(grid), len(grid[0])

	var ans int = 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '.' {
				continue
			}
			cnt := 0
			for _, dir := range dirs {
				nr, nc := r + dir[0], c + dir[1] 
				if !valid(nr, nc) {
					continue
				}

				if grid[nr][nc] == '@' {
					cnt++
				}
			}
			if (cnt < 4) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func valid(r, c int) bool {
	return r >= 0 && c >= 0 && r < rows && c < cols 
}