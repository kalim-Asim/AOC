package part

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	in every column from ending, get all element
	from top to bottom, add in a string
	if reached last row store that no.
	if a col has no no. perform that operation
*/

func Part2(lines []string) {
	var ans, rows, maxCol int = 0, len(lines), 0
	var op []string = strings.Fields(lines[rows-1])

	for _, line := range lines {
		maxCol = max(maxCol, len(line)) 
	}

	var nums [][]string 
	var num []string 
	for col := maxCol - 1; col >= 0; col-- {
		var found bool = false
		var s string 
		
		for row := 0; row < rows - 1; row++ {
			if lines[row][col] == ' ' {
				continue
			}
			found = true 
			s += string(lines[row][col])
		}
		if found {
			num = append(num, s)
		} 

		if !found || col == 0 {
			/*
	THIS IS WRONG !
	nums = append(nums, num)
	num = num[:0]
	CREATES REFERENCE TO SLICE, MODIFIES NUMS!!
			*/

			_copy := make([]string, len(num))
			copy(_copy, num)
			nums = append(nums, _copy)
			num = num[:0] // safe reuse
		}
	}
	reverse(nums)
	
	for i, x := range op {
		var sum int
		if x == "*" { sum = 1 }

		for _, s := range nums[i] {
			el, _ := strconv.Atoi(s)
			if x == "*" { 
				sum *= el
			} else {
				sum += el 
			}
		}

		ans += sum 
	}
	fmt.Println(ans)
}

func reverse(nums [][]string) {
	for i, j := 0, len(nums)-1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}