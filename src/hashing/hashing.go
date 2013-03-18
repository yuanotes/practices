package hashing

type Element struct {
    value string
    next *Element
}

// linked list hashtable
type HashTable struct {
    table_size int
    alphabet_size int
    table []*Element
}

// open addressing hashtable
type HashArray struct {
    table_size int
    alphabet_size int
    table []string
}

func (h *HashArray) Lookup(s string) bool {
    index := h.Hash(s)
    if index < 0 || index >= h.table_size {
        return false
    }
    ele := h.table[index]
    if ele != "" {
        if ele == s { return true } // found at the hash index
        original_index := index
        for {
            index = (index + 1) % h.table_size
            if index == original_index { return false } // avoid infinite loop
            if h.table[index] == s { return true }
        }
    }
    return false
}

func (h *HashArray) Insert(s string) bool {
    index := h.Hash(s)
    if index < 0 || index >= h.table_size {
        return false
    }
    if h.table[index]== "" {
        h.table[index] = s
    } else {
        original_index := index
        for {
            index = (index + 1) % h.table_size
            if index == original_index { return false } // avoid infinite loop
            if h.table[index] == "" {
                h.table[index] = s
                break
            }
        }
    }
    return true
}

func (h *HashArray) Hash(s string) int {
    result := 0
    for i, v := range s {
        result += h.alphabet_size^i * int(v)
    }
    return result % h.table_size
}

func (h *HashArray) Delete(s string) bool {
    index := h.Hash(s)
    if index < 0 && index >= h.table_size { return false }
    if h.table[index] == "" { return false }
    if h.table[index] == s {
        h.table[index] = "<DELETED>"
        return true
    }
    original_index := index
    for {
        index = (index + 1) % h.table_size
        if index == original_index { return false } // avoid infinite loop
        if h.table[index] == s {
            h.table[index] = "<DELETED"
            return true
        }
    }
    return false
}

func (h *HashTable) Lookup(s string) *Element {
    index := h.Hash(s)
    if index < 0 || index >= h.table_size {
        return nil
    }
    for ele := h.table[index]; ele != nil; ele=ele.next {
        if ele.value == s {
            return ele
        }
    }
    return nil
}

// Return a number from string.
func (h *HashTable) Hash(s string) (int) {
    result := 0
    for i, v := range s {
        result += h.alphabet_size^i * int(v)
    }
    return result % h.table_size
}

func (h *HashTable) Insert(s string) *Element {
    index := h.Hash(s)
    if index < 0 || index >= h.table_size {
        return nil
    }
    ele := h.table[index]
    if ele != nil {
        for {
            if ele.next == nil {
                ele.next = &Element{s, nil}
                break
            }
        }
        ele = &Element{s, nil}
    } else {
        h.table[index] = &Element{s, nil}
    }
    return ele
}

func (h *HashTable) Delete(s string) bool {
    index := h.Hash(s)
    if index < 0 && index >= h.table_size { return false }
    ele := h.table[index]
    if ele.value == s {
        h.table[index] = ele.next
        return true
    }
    for ; ele != nil ; ele=ele.next{
        if ele.value == s {
            ele = ele.next
            return true
        }
    }
    return false
}
