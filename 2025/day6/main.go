package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	var op []string
	var nums [][]int

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		data := strings.Fields(line)
		if data[0] == "*" || data[1] == "+" {
			op = data
		} else {
			arr := make([]int, len(data))
			for i := 0; i < len(data); i++ {
				num, _ := strconv.Atoi(data[i])
				arr[i] = num
			}
			nums = append(nums, arr)
		}
	}
	part1(op, nums)

	if err := sc.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}
}

func part1(op []string, nums [][]int) {
	var ans int

	for c := range nums {
		num, mul := 0, false 
		if op[c] == "*" {
			mul = true
			num = 1
		} 
		for r := range nums {
			if (mul) {
				num *= nums[r][c]
			} else {
				num += nums[r][c]
			}
		}
		ans += num
	}

	fmt.Println("Part1 :", ans)
}

/*
func part2(op []string, nums [][]int) {
	var ans int 

}
*/
