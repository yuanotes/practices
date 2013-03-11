package main

import (
    "utils"
    "fmt"
)

func insertion_sort(array []int) {
    for j, _ := range array {
        key := array[j]
        i := j-1
        for ; i>=0 && array[i] > key; i-=1 {
            array[i+1] = array[i]
        }
        array[i+1] = key
    }
}

func main(){
    array_1 := utils.GetIntSlice(10)
    array_2 := utils.GetIntSlice(10000)
    utils.PrintRunningTime(insertion_sort, array_1)
    utils.PrintRunningTime(insertion_sort, array_2)
    fmt.Println(array_1)
    fmt.Println(array_2)
}
