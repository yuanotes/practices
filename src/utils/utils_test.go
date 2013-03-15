package utils

import (
    "testing"
)

func TestGetRandInt(t *testing.T) {
    for i :=0; i < 1000; i++ {
        l := 1
        u := 4
        v := GetRandInt(l, u)
        if v < l || v >=u {
            t.Errorf("GetRandInt product out range number. l: %d, u: %d, v: %d.\n", l, u, v)
        }
    }
}
