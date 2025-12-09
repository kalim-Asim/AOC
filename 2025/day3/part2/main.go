package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

var k int = 12
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ans int 
	for scanner.Scan() {
		data := scanner.Text()
		if data == "" {
			continue
		}

		str, start := "", 0
		for i := 0; i < k; i++ {
			end := len(data) - (k - (i + 1))
      ch, newStart := find(start, data, end)
      str += string(ch)
			start = newStart
		}
		num, _ := strconv.Atoi(str)
		ans += num
	}
	fmt.Println(ans)
}

func find(l int, s string, r int)(rune, int) {
	index, largest := l, rune(s[l])
	for i := l; i < r; i++ {
		if rune(s[i]) > largest {
			largest = rune(s[i])
			index = i 
		}
	}
  return largest, index + 1
}