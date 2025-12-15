package main

import (
	"bufio"
	"day11/utils"
	"day11/parts"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close() 

	sc := bufio.NewScanner(f)
	var graph []utils.Edge

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		nodes := strings.Split(line, ": ")

		from := nodes[0]
		tos := (strings.Split(nodes[1], " "))
		to := []string{}
		for _, v := range tos {
			to = append(to, v)
		}

		edge := utils.Edge{
			From: from, 
			To: to,
		}

		graph = append(graph, edge)
	}

	parts.Part1(graph)
	parts.Part2(graph)
}