package main

import (
    "fmt"
)

func TrimSpace(s []byte) []byte{
    b := s[:0] // 利用0长切片
    for _, v := range s {
        fmt.Printf("this char: %c\n", v)
        if v != ' ' {
            b = append(b, v)
        }
    }
    return b
}

func main() {
    b := []byte{'g', ' ', 'l', ' ', 'n', 'g'}
    fmt.Printf("Before trim: %s, len: %d, cap: %d\n", b, len(b), cap(b))
    new_b := TrimSpace(b)
    fmt.Printf("After trim: %s, len: %d, cap: %d\n", new_b, len(new_b), cap(new_b))
}
