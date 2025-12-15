package parts

import (
	"day10/utils"
	"fmt"
)

var dp map[uint64]int
func Part2(data []utils.Data) {
	dp = make(map[uint64]int)

	n = len(data)
	var ans int 

	for _, d := range data {
		ans += f2(0, make([]int, len(d.Joltage)), d.Joltage, d.Buttons, len(d.Joltage), len(d.Buttons))
	}

	fmt.Println("Part2: ", ans)
}

func f2(index int, current []int, target []int, arr [][]int, bulbs int, buttons int) int {
	if index == buttons {
		if equal(current, target) {
			return 0
		}
		return INF
	}
	if exceeded(bulbs, current, target) {
		return INF 
	}

	key := utils.MakeKey(index, current)

	if val, ok := dp[key]; ok {
			return val
	}

	res := INF
	limit := maxPress(current, target, arr[index])

	for cnt := 0; cnt <= limit; cnt++ {
		if cnt >= res {
			break
		}
		
		next := append([]int(nil), current...)
		for _, i := range arr[index] {
				next[i] += cnt
		}
		res = min(res, cnt + f2(index+1, next, target, arr, bulbs, buttons))
	}
	dp[key] = res
	return dp[key]
}

func equal(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// if current exceeds the target value
func exceeded(tot int, curr []int, targ []int) bool {
	for i := 0; i < tot; i++ {
		if curr[i] > targ[i] {
			return true  
		}
	}
	return false
}

// max cnt each button can take
func maxPress(curr, targ []int, indices []int) int {
    limit := INF
    for _, i := range indices {
        if curr[i] < targ[i] {
            limit = min(limit, targ[i]-curr[i])
        }
    }
    return limit
}
