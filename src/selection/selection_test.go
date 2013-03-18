package selection

import (
    "testing"
    "utils"
    "sort"
    "fmt"
)


func TestSelect1(t *testing.T) {
    count := 20
    array := utils.GetIntSlice(count, false)
    fmt.Println("Initialized array:", array)
    k := utils.GetRandInt(0, count)
    Select1(0, len(array)-1, k, array)
    fmt.Println(array)
    select_v := array[k]
    sort.Ints(array)
    fmt.Println(array)
    fmt.Println("k is", k)
    real_v := array[k]
    if select_v != real_v {
        t.Errorf("The kth smallest number is not correct. select: %d, real: %d", select_v, real_v)
    }
}
