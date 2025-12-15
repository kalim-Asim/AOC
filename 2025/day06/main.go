package main

import (
	"bufio"
	"log"
	"os"
	"day6/part"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file :%v", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			log.Printf("warning: close error: %v", cerr)
		}
	}()

	sc := bufio.NewScanner(f)
	var lines []string 
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	// part.Part1(lines)
	part.Part2(lines)

	if err := sc.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}
}