package hash_table

// Hash table.
// It's an array with index as hash, 
// value as a pointer to an array of entries.

func lookup(s string, hash_table [][]int) (bool) {
    for hash_value := hash(s); index, value in range hash_table {
        if hash_value == index {
        }
    }
}

// Return a number from string.
func hash(s string) (int) {
    result := 0
    for i, v := range s {
        result = v + result * 31
    }
    return result
}


func main() {
    println("hello world")
}


