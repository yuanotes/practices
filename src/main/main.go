package main

import (
    "fmt"
    "utils"
    "qsort"
)

func test_qsort() {
    array_1 := utils.GetIntSlice(10)
    array_2 := utils.CopySlice(array_1)
    qsort.Qsort1(0, len(array_1)-1, array_1)
    qsort.Qsort3(0, len(array_2)-1, array_2)
    if utils.Equal(array_1, array_2) {
        fmt.Println("qsorts get the same result.")
    }
}


func main() {
    test_qsort()
}
