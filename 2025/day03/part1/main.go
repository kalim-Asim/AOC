package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ans int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		
		idx1, first := 0, 0
		for i := 0; i < len(s) - 1; i++ {
			if int(s[i] - '0') > first {
				first = int(s[i] - '0')
				idx1 = i 
			}
		}

		sec := 0
		for i := idx1 + 1; i < len(s); i++ {
			if int(s[i] - '0') > sec {
				sec = int(s[i] - '0')
			}
		}

		ans += (first * 10 + sec)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(ans)
}