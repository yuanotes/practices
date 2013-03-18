package hashing

type Element struct {
    value string
    next *Element
}

func Lookup(s string, hash_table []*Element) *Element {
    index :=Hash(s)
    if hash_table[index] != nil {
        ele := hash_table[index]
        for ; ele.next == nil; ele=ele.next {
            if ele.value == s {
                return ele
            }
        }
        ele.next = &Element{s, nil}
    } else {
        hash_table[index] = &Element{s, nil} 
    }
    return nil
}

// Return a number from string.
func Hash(s string) (int) {
    result := 0
    for i, v := range s {
        result += 26^i * int(v)
    }
    return result % 255
}

func Init() []*Element {
    hash_table := make([]*Element, 255)
    return hash_table
}
