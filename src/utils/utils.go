package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

func GetRandInt(low, up int) int {
    // Get a random int between [low, up)
    if low >= up {
        return up
    }
    rand.Seed(time.Now().UTC().UnixNano())
    count := up - low
    value := rand.Intn(count)
    return low + value
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, c := range a {
        if c != b[i] {
            return false
        }
    }
    return true
}

func CopySlice(a []int) []int {
    b := make([]int, len(a), cap(a))
    copy(b, a)
    return b
}

func GetIntSlice(length int, equal bool) []int {
	rand.Seed(time.Now().UTC().UnixNano())
    var array []int
    if equal {
        array := make([]int, length)
        value := rand.Int()
        for i,_ := range array {
            array[i] = value
        }
    } else {
        array = rand.Perm(length)
    }
	return array
}

func PrintRunningTime(fun interface{}, params ...interface{}) {
	fun_name := runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name()
	function := reflect.ValueOf(fun)
	args := []reflect.Value{}
	if params != nil {
		args = make([]reflect.Value, len(params))
		for i, value := range params {
			args[i] = reflect.ValueOf(value)
		}
	}
	ts := time.Now().UnixNano()
	function.Call(args)
	te := time.Now().UnixNano()
	fmt.Printf("Total running time of %s is %g micro seconds.\n", fun_name, float64(te-ts)/1000000)
}
