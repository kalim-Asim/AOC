package main

import (
	"bufio"
	"day10/parts"
	. "day10/utils"
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
	var lines []Data 
	for sc.Scan() {
		line := strings.TrimSpace(string(sc.Text()))
		light, button, joltage := splitParts(line)
		data := Data{
			Lights: light,
			Buttons: button,
			Joltage: joltage,
		}
		lines = append(lines, data)
	}

	parts.Part1(lines)
	parts.Part2(lines)
}

func splitParts(s string) ([]int, [][]int, []int) {
	var arr2 [][]int 
	var arr1, arr3 []int
	i := 0
	for i < len(s) {
		switch s[i] {
		case '[':
			j := strings.IndexByte(s[i:], ']')
			sq := s[i+1 : i+j]
			for k := 0; k < len(sq); k++ {
				if sq[k] == '#' {
					arr1 = append(arr1, k)
				}
			}
			i += j + 1

		case '(':
			j := strings.IndexByte(s[i:], ')')
			arr := strings.Split(s[i+1:i+j], ",")
			nums := []int{}
			for _, v := range arr {
				num, _ := strconv.Atoi(v)
				nums = append(nums, num)
			}
			arr2 = append(arr2, nums)
			i += j + 1

		case '{':
			j := strings.IndexByte(s[i:], '}')
			curly := s[i+1 : i+j]
			i += j + 1
			str := strings.Split(curly, ",")
			for _, a := range str {
				num, _ := strconv.Atoi(a)
				arr3 = append(arr3, num)
			}
		default:
			i++
		}
	}

	return arr1, arr2, arr3
}