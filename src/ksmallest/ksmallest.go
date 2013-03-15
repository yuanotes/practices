package ksmallest

import (
    "utils"
)

func Select1(l, u, k int, array []int) {
    //Little modification to qsort 4.
    //Just make sure that kth element is in correct position.
    if l >= u {
        return
    }
    rand_index := utils.GetRandInt(l, u+1)
    array[rand_index], array[l] = array[l], array[rand_index]

    t :=array[l]
    i :=l + 1
    j :=u

    for {
        for ;i<=u && array[i] < t; i++ {}
        for ;array[j] > t; j-- {}
        if i > j {
            break
        }
        array[i], array[j] = array[j], array[i]
    }
    array[l], array[j] = array[j], array[l]
    if j < k {
        Select1(j+1, u, k, array)
    } else {
        Select1(l, j-1, k, array)
    }
}
