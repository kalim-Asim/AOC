package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			log.Printf("warning: close error: %v", cerr)
		}
	}()

	sc := bufio.NewScanner(f)
	part1(sc)
	part2(sc)

	// check for scanning errors (non-EOF)
	if err := sc.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}
}


func part1(sc *bufio.Scanner) {
	ans := 0
	var ranges [][2]int

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			break
		}
		lr := strings.Split(line, "-")
		l, _ := strconv.Atoi(lr[0])
		r, _ := strconv.Atoi(lr[1])
		ranges = append(ranges, [2]int{l, r})
	}

	sort.Slice(ranges, func(i, j int) bool{
		return ranges[i][0] < ranges[j][0]
	})

	merged := make([][2]int, 0)

	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1][1] < r[0] - 1 {
			merged = append(merged, r)
		} else {
			if r[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = r[1]
			}
		}
	}


	for sc.Scan() {
		s := strings.TrimSpace(sc.Text())
		if s == "" {
			continue
		}
		num, _ := strconv.Atoi(s)

		// Binary search to find interval with start <= x
		idx := sort.Search(len(merged), func(i int) bool {
				return merged[i][0] > num
		})
		// idx is first interval which has L > x
		// idx--, last interval which has L <= x
		if idx == 0 {
			continue
		}
		idx--
		if merged[idx][0] <= num && num <= merged[idx][1] {
			ans++
		}
	}
	fmt.Println("Part 1:", ans)
}

func part2(sc *bufio.Scanner)  {
	ans := 0
	var ranges [][2]int
	var interval = true
	for sc.Scan() {
		if !interval {
			break
		}
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			interval = false
			continue
		}

		lr := strings.Split(line, "-")
		l, _ := strconv.Atoi(lr[0])
		r, _ := strconv.Atoi(lr[1])
		ranges = append(ranges, [2]int{l, r})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	
	merged := make([][2]int, 0)
	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1][1] < r[0]-1 {
			merged = append(merged, r)
		} else {
			if r[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = r[1]
			}
		}
	}

	for _, r := range merged {
		ans += r[1] - r[0] + 1
	}
	fmt.Println("Part 2:", ans)
}
