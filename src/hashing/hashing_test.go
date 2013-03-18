package hashing

import (
    "testing"
)


func TestHash(t *testing.T) {
    str1 := "hello world"
    if Hash(str1) == 0 {
        t.Errorf("Hash function get 0...")
    }
}

func TestLookup(t *testing.T) {
    str1 := "hello world"
    str2 := "你好世界"
    str3 := ""
    hash_table := Init()
    if Lookup(str1, hash_table) != nil {
        t.Error("The string should be non-exist.")
    }
    if Lookup(str2, hash_table) != nil {
        t.Error("The string should be non-exist.")
    }
    if Lookup(str3, hash_table) != nil {
        t.Error("The string should be non-exist.")
    }
    if Lookup(str1, hash_table) == nil {
        t.Errorf("The string %v should be exist in %v", str1, hash_table)
    }
    if Lookup(str2, hash_table) == nil {
        t.Errorf("The string %v should be exist in %v", str2, hash_table)
    }
    if Lookup(str3, hash_table) == nil {
        t.Errorf("The string %v should be exist in %v", str3, hash_table)
    }
    t.Logf("hash_table is %v", hash_table)
    
}
