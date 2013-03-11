package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

func Swap(a, b int, array []int) {
	// can't use xor when a == b
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}

func GetIntSlice(length int) []int {
	array := make([]int, length)
	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range array {
		array[i] = rand.Intn(100)
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
