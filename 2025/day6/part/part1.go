package part

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) {
	var ans int
	var op []string
	var nums [][]int

	for _, line := range lines {
		data := strings.Fields(line)

		// operator row
		if data[0] == "*" || data[0] == "+" {
			op = data
			continue
		}

		// number rows
		arr := make([]int, len(data))
		for i := range data {
			num, _ := strconv.Atoi(data[i])
			arr[i] = num
		}
		nums = append(nums, arr)
	}

	// compute
	for c := range nums[0] {
		num, mul := 0, false

		if op[c] == "*" {
			mul = true
			num = 1
		}

		for r := range nums {
			if mul {
				num *= nums[r][c]
			} else {
				num += nums[r][c]
			}
		}

		ans += num
	}

	fmt.Println("Part1:", ans)
}
