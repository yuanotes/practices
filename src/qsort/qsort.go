package qsort

func Qsort1(l, u int, array []int) {
	if l >= u {
		return
	}
	m := l
	// find the position of array[l]
	for i := l + 1; i <= u; i ++ {
		if array[i] < array[l] {
			m += 1
            array[m], array[i] = array[i], array[m]
		}
	}
    array[m], array[l] = array[l], array[m]
	Qsort1(l, m-1, array)
	Qsort1(m+1, u, array)
}

func Qsort3(l, u int, array []int) {
    // Partition on both side.
    if l>=u {
        return
    }
    t := array[l]
    i := l + 1;
    j := u;
    for {
        if i<=u && array[i] < t {
            i += 1
        }
        if array[j] > t {
            j -= 1
        }
        if i > j {
            break
        }
        array[i], array[j] = array[j], array[i]
    }
    array[l], array[j] = array[j], array[l]
    Qsort3(l, j-1, array)
    Qsort3(j+1, u, array)
}
