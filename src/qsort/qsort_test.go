package qsort

import (
    "testing"
    "sort"
    "utils"
)

func TestQsort1(t *testing.T) {
    var a []int
    // test equality
    a = utils.GetIntSlice(10000, true)
    Qsort1(0, len(a)-1, a)
    if !sort.IntsAreSorted(a) {
        t.Errorf("Qsort1 error on equal slice.")
    }

    // test sorting
    a = utils.GetIntSlice(10000, false)
    Qsort1(0, len(a)-1, a)
    if !sort.IntsAreSorted(a) {
        t.Errorf("Qsort1 error on non-equal slice.")
    }
}

func BenchmarkQsort1(b *testing.B) {
    b.StopTimer()
    a := utils.GetIntSlice(10, false)
    b.StartTimer()
    for i:=0; i<b.N; i++ {
        Qsort1(0, len(a)-1, a)
    }
}

func BenchmarkQsort3(b *testing.B) {
    b.StopTimer()
    a := utils.GetIntSlice(10, false)
    b.StartTimer()
    for i:=0; i<b.N; i++ {
        Qsort3(0, len(a)-1, a)
    }
}


func TestQsort3(t *testing.T) {
    var a []int
    // test equality
    a = utils.GetIntSlice(10000, true)
    Qsort3(0, len(a)-1, a)
    if !sort.IntsAreSorted(a) {
        t.Errorf("Qsort1 error on equal slice.")
    }

    // test sorting
    a = utils.GetIntSlice(10000, false)
    Qsort3(0, len(a)-1, a)
    if !sort.IntsAreSorted(a) {
        t.Errorf("Qsort1 error on non-equal slice.")
    }
}



