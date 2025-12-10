package main

import (
	"fmt"
	"strings"
	"os"
)

var rows, cols, ans int
var dirs = [][]int{
	{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	{-1, 1}, {-1, -1}, {1, -1}, {1, 1},
}
func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	raw := strings.Split(strings.TrimSpace(string(file)), "\n")

	grid := make([][]byte, len(raw))
	for i := range raw {
			line := strings.TrimRight(raw[i], "\r")
			grid[i] = []byte(line)     // convert to byte slice
	}

	rows, cols = len(grid), len(grid[0])

	var flag bool = true 

	for flag {
		flag = false
		mark := make([][]bool, rows)
		for i := range mark {
			mark[i] = make([]bool, cols)
		}

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == '.' {
					continue
				}
				cnt := 0

				for _, dir := range dirs {
					nr, nc := r + dir[0], c + dir[1]
					if !valid(nr, nc) || grid[nr][nc] == '.' {
						continue
					}
					cnt++
				}

				if cnt < 4 {
					flag = true
					ans++
					mark[r][c] = true 
				}
			}
		}

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if mark[r][c] {
					grid[r][c] = '.'
				}
			} 
		}
	}
	
	fmt.Println(ans)
}

func valid(nr, nc int) bool {
	return nr >= 0 && nc >=0 && nr < rows && nc < cols
}