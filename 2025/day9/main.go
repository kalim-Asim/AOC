package main

import (
	"bufio"
	"day9/parts"
	. "day9/utils"
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
		x, y := toInt(strings.Split(line, ",")[0]), toInt(strings.Split(line, ",")[1])
		points = append(points, Point{x, y})
	}

	parts.Part1(points)
	parts.Part2(points)
}

func toInt(s string) int {
    n, _ := strconv.Atoi(s)
    return n
}




