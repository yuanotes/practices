package heapsort

import (
    "testing"
    "utils"
    "sort"
)

func TestCheckPriority(t *testing.T) {
    array := []int{}
    if !CheckPriority(array, false) {
        t.Errorf("CheckPriority cannot check empty array.")
    }
    array = []int{10}
    if !CheckPriority(array, false) {
        t.Errorf("CheckPriority cannot check one-element array.")
    }
    array = []int{1,2,3,4,5,6,7}
    if !CheckPriority(array, false) {
        t.Errorf("CheckPriority cannot check multiple elements array.")
    }
    array = []int{8,7,6,5,4,3,2,1}
    if !CheckPriority(array, true) {
        t.Errorf("CheckPriority cannot check multiple elements array.")
    }
    array = []int{8,2,3,4,5,6,7}
    if CheckPriority(array, false) {
        t.Errorf("CheckPriority cannot check multiple elements array.")
    }
    array = []int{8,7,6,5,4,3,4}
    if CheckPriority(array, false) {
        t.Errorf("CheckPriority cannot check multiple elements array.")
    }
}

func TestSiftup(t *testing.T) {
    var array []int
    for i, _:= range array {
        Siftup(array[:i+1])
    }
    if !CheckPriority(array, false) {
        t.Errorf("Siftup corrupt the priority of empty array.")
    }
    array = []int{1}
    for i, _:= range array {
        Siftup(array[:i+1])
    }
    if !CheckPriority(array, false) {
        t.Errorf("Siftup corrupt the priority of one-element array.")
    }
    array = utils.GetIntSlice(1000000, false)
    for i, _:= range array {
        Siftup(array[:i+1])
    }
    if !CheckPriority(array, false) {
        t.Errorf("Siftup corrupt the priority of multi-element array.")
    }
}

func TestSiftupInverse(t *testing.T) {
    var array []int
    for i, _:= range array {
        SiftupInverse(array[:i+1])
    }
    if !CheckPriority(array, true) {
        t.Errorf("SiftupInverse corrupt the priority of empty array.")
    }
    array = []int{1}
    for i, _:= range array {
        SiftupInverse(array[:i+1])
    }
    if !CheckPriority(array, true) {
        t.Errorf("SiftupInverse corrupt the priority of one-element array.")
    }
    array = utils.GetIntSlice(1000000, false)
    for i, _:= range array {
        SiftupInverse(array[:i+1])
    }
    if !CheckPriority(array, true) {
        t.Errorf("SiftupInverse corrupt the priority of multi-element array.")
    }
}

func TestSiftdown(t *testing.T) {
    var array []int
    if Siftdown(array) != -1 {
        t.Errorf("Siftdown cannot manipulate an empty array.")
    }
    array = []int{1}
    if Siftdown(array) != 1 {
        t.Errorf("Siftdown cannot manipulate a one-element array.")
    }
    array = utils.GetIntSlice(1000000, false)
    for i, _:= range array {
        Siftup(array[:i+1])
    }
    get_min := Siftdown(array)
    sort.Ints(array)
    real_min := array[0]
    if get_min != real_min {
        t.Errorf("Siftdown doesn't get the minimum number.")
    }
    if !CheckPriority(array, false) {
        t.Errorf("Siftdown corrupt the priority of the array.")
    }
}

func CheckPriority(array []int, reverse bool) bool {
    n := len(array) - 1
    if n < 1 { // return true for empty and one-element array.
        return true
    }
    for i, _ := range array {
        c := (i + 1) * 2 - 1 // left child
        if c <= n {
            if reverse {
                if c + 1 <= n && array[c+1] < array[c] {
                    c ++
                }
                if array[i] < array[c] {
                    return false
                }
            } else {
                if c + 1 <= n && array[c+1] > array[c] {
                    c ++
                }
                if array[i] > array[c] {
                    return false
                }
            }
        }
    }
    return true
}

func TestPorisort(t *testing.T) {
    array := utils.GetIntSlice(1000000, false)
    Priosort(array)
    if !sort.IntsAreSorted(array) {
        t.Errorf("Porioty queue is not working...")
    }
}

func TestHeapsort(t *testing.T) {
    array := utils.GetIntSlice(1000000, false)
    Heapsort(array)
    if !sort.IntsAreSorted(array) {
        t.Errorf("Heap sort can't sort a randomized array...")
    }
}
