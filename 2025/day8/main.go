package main

import (
	"bufio"
	"day8/parts"
	. "day8/utils" 
	"os"
	"strconv"
	"strings"
)


func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	points := make([]Point, 0)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		values := strings.Split(line, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])

		point := Point{
			X: x, 
			Y: y,
			Z: z,
		}

		points = append(points, point)
	}
	
	parts.Part1(points)
	parts.Part2(points)
} 
