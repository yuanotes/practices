package heapsort

func Siftdown(array []int) int {
    n := len(array) - 1
    if n < 0 { return -1 } // except empty array
    if n == 0 { return array[n] } // return one-element array
    pop := array[0]
    i := 0
    array[0] = array[n]
    array[n] = 0 //empty the last one
    for {
        c := (i+1) * 2 - 1 // left child of i
        if c > n - 1 {
            break
        }
        if c + 1 <=n - 1 && array[c] > array[c+1] {
                c++ //right child of i
        }
        if array[i] <= array[c] {
            break
        }
        array[i], array[c] = array[c], array[i]
        i = c
    }
    return pop
}


func SiftdownInverse(array []int) int {
    n := len(array) - 1
    if n < 0 { return -1 } // except empty array
    if n == 0 { return array[n] } // return one-element array
    pop := array[0]
    i := 0
    array[0] = array[n]
    array = array[:n] //empty the last one
    for {
        c := (i+1) * 2 - 1 // left child of i
        if c > n - 1 {
            break
        }
        if c + 1 <=n - 1 && array[c] < array[c+1] {
                c++ //right child of i
        }
        if array[i] >= array[c] {
            break
        }
        array[i], array[c] = array[c], array[i]
        i = c
    }
    return pop
}

func Siftup(array []int) {
    n := len(array) - 1 // the last element
    for {
        if n <= 0 { // the array is empty or just one value and the end
            return
        }
        p := (n-1)/2 // get parent position of array[n]
        if array[n] >= array[p] {
            break
        } else {
            array[n], array[p] = array[p], array[n]
            n = p
        }
    }
}

func SiftupInverse(array []int) {
    n := len(array) - 1 // the last element
    for {
        if n <= 0 { // the array is empty or just one value
            return
        }
        p := (n-1)/2 // get parent position of array[n]
        if array[n] <= array[p] {
            break
        } else {
            array[n], array[p] = array[p], array[n]
            n = p
        }
    }
}

type PrioQueue struct {
    array []int
}

func (q *PrioQueue) Insert(value int) {
    q.array = append(q.array, value) // this may cost a lot
    Siftup(q.array)
}

func (q *PrioQueue) Exactmin() int {
    min := Siftdown(q.array)
    q.array = q.array[:len(q.array)-1]
    return min
}

func Priosort(array []int) {
    // A basic sort function that builds
    // a heap for an array and exact element
    // from the array.
    q := &PrioQueue{}
    for _, v := range array {
        q.Insert(v)
    }
    for i, _ := range array {
        array[i] = q.Exactmin()
    }
}

func Heapsort(array []int) {
    // The Porisort using an auxiliary array
    // for sorting. While Heapsort will sort the 
    // array in place.
    for i, _:= range array {
        SiftupInverse(array[:i+1])
    }
    for n:=len(array) -1 ; n >=0 ; n -- {
        array[n] = SiftdownInverse(array[:n+1])
    }
}
