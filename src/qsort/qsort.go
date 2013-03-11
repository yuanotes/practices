package main

import (
	//"fmt"
    "utils"
)

func qsort1(l, u int, array []int) {
	if l >= u {
		return
	}
	m := l
	// find the position of array[l]
	for i := l + 1; i <= u; i++ {
		if array[i] < array[l] {
			m += 1
			utils.Swap(m, i, array)
		}
	}
	utils.Swap(m, l, array)
	qsort1(l, m-1, array)
	qsort1(m+1, u, array)
}

func main() {
	array := utils.GetIntSlice(100000)
	utils.PrintRunningTime(qsort1, 0, len(array)-1, array)
    //fmt.Println(array)
}
