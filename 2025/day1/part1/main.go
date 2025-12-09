package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)

  var ans, cur int = 0, 50
	for scanner.Scan() {
		line := scanner.Text() 

		if line == "" {
			continue
		} 

		dir := rune(line[0])
		num, _ := strconv.Atoi(line[1:])

		if (dir == 'L') {
			cur = (cur - num + 100) % 100
		} else {
			cur = (cur + num) % 100
		} 

		if cur == 0 { ans++ }
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}	

	fmt.Println(ans)
}

