package hashing

import (
    "testing"
)

var h = &HashTable{255, 26, make([]*Element, 255)}
var h2 = &HashArray{255, 26, make([]string,255)}

func TestLinkedlistHashTable(t *testing.T) {
    str1 := "hello world"
    str2 := "你好世界"
    str3 := ""
    if h.Hash(str1) == 0 {
        t.Errorf("Hash function get 0...")
    }
    if h.Hash(str2) == 0 {
        t.Errorf("Hash function get 0...")
    }
    if h.Hash(str3) != 0 {
        t.Errorf("Hash function get 0...")
    }
    if h.Lookup(str1) != nil {
        t.Error("The string should be non-exist.")
    }
    if h.Lookup(str2) != nil {
        t.Error("The string should be non-exist.")
    }
    if h.Lookup(str3) != nil {
        t.Error("The string should be non-exist.")
    }
    h.Insert(str1)
    h.Insert(str2)
    h.Insert(str3)
    if h.Lookup(str1) == nil {
        t.Errorf("The string should exist in table.")
    }
    if h.Lookup(str2) == nil {
        t.Errorf("The string should exist in table.")
    }
    if h.Lookup(str3) == nil {
        t.Errorf("The string should exist in table.")
    }
    h.Insert(str1)
    h.Delete(str1)
    h.Delete(str1)
    h.Delete(str2)
    h.Delete(str3)
    if h.Lookup(str1) != nil {
        t.Error("The string should be exist.")
    }
    if h.Lookup(str2) != nil {
        t.Error("The string should be non-exist.")
    }
    if h.Lookup(str3) != nil {
        t.Error("The string should be non-exist.")
    }
}



func TestHashArray(t *testing.T) {
    str1 := "hello world"
    str2 := "你好世界"
    str3 := "!@#@@#$%#"
    if h2.Hash(str1) == 0 {
        t.Errorf("Hash function get 0...")
    }
    if h2.Hash(str2) == 0 {
        t.Errorf("Hash function get 0...")
    }
    if h2.Hash(str3) == 0 {
        t.Errorf("Hash function get 0...")
    }
    if h2.Lookup(str1) {
        t.Error("The string should be non-exist.")
    }
    if h2.Lookup(str2) {
        t.Error("The string should be non-exist.")
    }
    if h2.Lookup(str3) {
        t.Error("The string should be non-exist.")
    }
    h2.Insert(str1)
    h2.Insert(str2)
    h2.Insert(str3)
    t.Logf("now hash array is %v", h2.table)
    if !h2.Lookup(str1) {
        t.Errorf("The string should exist in table.")
    }
    if !h2.Lookup(str2) {
        t.Errorf("The string should exist in table.")
    }
    if !h2.Lookup(str3) {
        t.Errorf("The string should exist in table.")
    }
    for i:=0;i<300;i++{
        h2.Insert(str1)
    }
    t.Log("now hash array is", h2.table)
    h2.Delete(str1)
    h2.Delete(str1)
    for i:=0;i<300;i++{
        h2.Delete(str1)
    }
    t.Log("now hash array is", h2.table)
    h2.Delete(str2)
    h2.Delete(str3)
    t.Log("now hash array is", h2.table)
    if h2.Lookup(str1) {
        t.Error("The string should be non-exist.")
    }
    if h2.Lookup(str2) {
        t.Error("The string should be non-exist.")
    }
    if h2.Lookup(str3) {
        t.Error("The string should be non-exist.")
    }
}
