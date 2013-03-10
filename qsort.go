package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

func swap(a, b int, array []int) {
	// can't use xor when a == b
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}

func get_int_array(length int) []int {
	array := make([]int, length)
	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range array {
		array[i] = rand.Intn(100)
	}
	return array
}

func qsort1(l, u int, array []int) {
	if l >= u {
		return
	}
	m := l
	// find the position of array[l]
	for i := l + 1; i <= u; i++ {
		if array[i] < array[l] {
			m += 1
			swap(m, i, array)
		}
	}
	swap(m, l, array)
	qsort1(l, m-1, array)
	qsort1(m+1, u, array)
}

func main() {
	array := get_int_array(1000000)
	get_running_time(qsort1, 0, len(array)-1, array)
	get_running_time(get_int_array, 100)
	fmt.Println(array)
}

func get_running_time(fun interface{}, params ...interface{}) {
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
