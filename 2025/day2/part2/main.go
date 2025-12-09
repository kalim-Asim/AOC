package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var N int = 100000
var arr []int 

func main() {
	precompute()

	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(file))
	intervals := strings.Split(data, ",")

	sum := 0

	for _, p := range intervals {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		lr := strings.Split(p, "-")
		if len(lr) != 2 {
			continue
		}
		l, _ := strconv.Atoi(lr[0])
		r, _ := strconv.Atoi(lr[1])
		
		sum += count(l, r)
	}

	fmt.Println(sum)
}

func removeDuplicates(arr []int) []int {
	mp := make(map[int]bool)
	for _, v := range arr {
		mp[v] = true 
	}

	res := make([]int, 0, len(mp))
	for k := range mp {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}
func precompute() {
	for i := 1; i <= N; i++ {
		s := strconv.Itoa(i)
		str_ := s 
		for len(s) <= 10 {
			s = s + str_ 
			if len(s) <= 10 {
				num, _ := strconv.Atoi(s)
				arr = append(arr, num)
			}
		}
	}
	arr = removeDuplicates(arr)
}
func lower_bound(x int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
			mid := (lo + hi) / 2
			if arr[mid] >= x {
					hi = mid
			} else {
					lo = mid + 1
			}
	}
	return lo
}
func upper_bound(x int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
			mid := (lo + hi) / 2
			if arr[mid] > x {
					hi = mid
			} else {
					lo = mid + 1
			}
	}
	return lo
}
func count(l int, r int) int {
	lb := lower_bound(l)
	ub := upper_bound(r)
	tot := 0
	for i := lb; i < ub; i++ {
		tot += arr[i]
	}
	return tot
}