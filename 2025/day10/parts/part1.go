package parts

import (
	. "day10/utils"
	"fmt"
)
/*
	what i am thinking? 
	- since we need to minimize no. of operation
	possible solution - recursion we memo
*/
var n int 
var newData []Data2
const INF = int(1e9)

func Part1(data []Data) {
	n = len(data)
	
	for _, v := range data {
		light, button := v.Lights, v.Buttons

		maskLight := 0 
		for _, l := range light {
			maskLight = maskLight | (1 << l)
		}
		var maskButtons []int 
		for _, butt := range button {
			maskButton := 0
			for _, b := range butt {
				maskButton = maskButton | (1 << b)
			}
			maskButtons = append(maskButtons, maskButton)
		}

		data2 := Data2{
			MaskLight: maskLight,
			MaskButtons: maskButtons,
		}
		newData = append(newData, data2)
	}

	var ans int 
	for _, v := range newData {
		ans += f(0, v.MaskLight, v.MaskButtons, 0)
	}

	fmt.Println("Part1: ", ans)
}

func f(i int, targetMask int, arr []int, currentMask int) int {
	if i == len(arr) {
		if targetMask == currentMask {
			return 0
		}
		return INF
	}

	notTake := f(i+1, targetMask, arr, currentMask)
	Take := 1 + f(i+1, targetMask, arr, currentMask ^ arr[i])

	return min(Take, notTake)
}