package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var N int = 100
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err) 
	} 
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ans, cur = 0, 50

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} 

		dir := rune(line[0])
		num, _ := strconv.Atoi(line[1:])
    
		ans += num / N
		num %= N
		if dir == 'L' {
			if cur != 0 && num > cur {
				ans++
			}
			cur = (cur - num + N) % N
		} else {
			if num > N - cur {
				ans++
			}
			cur = (cur + num) % N	
		}

		if cur == 0 {
			ans++
		}
	}
	
	if err := scanner.Err(); err != nil {
		panic(err)
	} 

	fmt.Println(ans)
} 
